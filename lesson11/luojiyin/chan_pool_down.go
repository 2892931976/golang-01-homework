package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func fetch(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	var urls []string

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")
		if ok {
			urls = append(urls, link)
		} else {
			fmt.Println("not found")
		}
	})
	return urls, nil
}

func cleanURL(taskq chan string, ur string, urls []string) {
	u, err := url.Parse(ur)
	if err != nil {
		fmt.Println("base url parse err")
	}
	scheme := u.Scheme
	host := u.Host
	p := strings.Split(u.Path, "/")
	path := strings.Join(p[:2], "")

	for _, link := range urls {
		//fmt.Println(link)
		switch {
		case strings.HasPrefix(link, "http") || strings.HasPrefix(link, "https"):
		case strings.HasPrefix(link, "//"):
			link = scheme + ":" + link
		case strings.HasPrefix(link, "/"):
			link = fmt.Sprintf("%s://%s%s", scheme, host, link)
			//taskq <- link
		default:
			//link = scheme + "://" + host + "/" + base_path + "/" + link
			link = fmt.Sprintf("%s://%s/%s%s", scheme, host, path, link)
		}
		//fmt.Println(link)
		taskq <- link
	}
}

func downimg(urls chan string, wg *sync.WaitGroup, tmpdir string) error {
	for link := range urls {
		reps, err := http.Get(link)
		if err != nil {
			fmt.Println("creating img file err ", err)
		}
		defer reps.Body.Close()

		s := path.Base(link)
		fullpath := filepath.Join(tmpdir, s)
		f, err := os.Create(fullpath)
		//fmt.Println(fullpath)
		defer f.Close()
		if err != nil {
			fmt.Println("get ulr error")
			continue
		}
		io.Copy(f, reps.Body)
	}
	wg.Done()
	return nil
}

func maketar(dir string, w io.Writer) error {
	baseDir := filepath.Base(dir)
	compress := gzip.NewWriter(w)
	defer compress.Close()
	tr := tar.NewWriter(compress)
	defer tr.Close()

	filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			fmt.Println("tar FileInfoHeader :", err)
			return err
		}

		p, _ := filepath.Rel(dir, name)
		header.Name = filepath.Join(baseDir, p)

		err = tr.WriteHeader(header)
		if err != nil {
			fmt.Println("tar WriteHeader err: ", err)
			return err
		}

		f, err := os.Open(name)
		if err != nil {
			fmt.Println("open file error:", err)
			return err
		}
		defer f.Close()

		io.Copy(tr, f)
		return nil

	})
	return nil
}

func main() {
	t1 := time.Now().Unix()

	url := "http://www.quanjing.com/"
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	tmpdir, err := ioutil.TempDir("", "img")
	if err != nil {
		log.Panic(err)
	}
	defer os.RemoveAll(tmpdir)

	taskq := make(chan string)

	wg := new(sync.WaitGroup)
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go downimg(taskq, wg, tmpdir)
	}
	cleanURL(taskq, url, urls)

	fmt.Println(tmpdir)

	close(taskq)
	wg.Wait()

	fmt.Println("sub time seconds >>>", int(time.Now().Unix()-int64(t1)))
	tarFile := filepath.Base(tmpdir)

	f, err := os.Create(tarFile + "tar.gz")
	if err != nil {
		log.Fatal(err)
	}

	if err := maketar(tmpdir, f); err != nil {
		log.Println(err)
	}

}
