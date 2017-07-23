package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"flag"
	"github.com/PuerkitoBio/goquery"
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
)

var (
	site  = flag.String("s", "http://daily.zhihu.com/", "url to download")
	label = flag.String("l", "img", "label to download")
	file  = flag.String("f", "img.tar.gz", "file to download")
	pool  = flag.Int("p", 5, "thread to download")
)

var labelAttr = map[string]string{
	"img":    "src",
	"script": "src",
	"a":      "href",
}

func checkUrl(url string) int {
	if strings.HasPrefix(url, "//") {
		return 1
	} else if strings.HasPrefix(url, "/") {
		return 2
	}
	if len(strings.Split(url, ":")) == 1 {
		return 3
	}
	return 0
}

func cleanUrl(u string) (string, error) {
	s, err := url.Parse(*site)
	if err != nil {
		return "", err
	}
	scheme, host, path := s.Scheme, s.Host, s.Path
	n := checkUrl(u)
	switch n {
	case 0:
	case 1:
		u = scheme + ":" + u
	case 2:
		u = scheme + "://" + host + u
	case 3:
		u = scheme + "://" + host + path + u
	}
	return u, nil
}

func fetchUrl(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find(*label).Each(func(i int, s *goquery.Selection) {
		l, ok := s.Attr(labelAttr[*label])
		if ok {
			urls = append(urls, l)
		}
	})
	return urls, nil
}

func downloadUrl(u string, dir string) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	path := filepath.Join(dir, path.Base(u))
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, resp.Body)
	return nil
}

func makeTar(dir string, w io.Writer) error {
	basedir := filepath.Base(dir)
	gw := gzip.NewWriter(w)
	defer gw.Close()
	tr := tar.NewWriter(gw)
	defer tr.Close()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		p, _ := filepath.Rel(dir, path)
		hdr.Name = filepath.Join(basedir, p)
		err = tr.WriteHeader(hdr)
		if err != nil {
			return err
		}
		fs, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fs.Close()
		if info.Mode().IsRegular() {
			io.Copy(tr, fs)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func work(ch chan string, wg *sync.WaitGroup, dir string) {
	for u := range ch {
		r, err := cleanUrl(u)
		if err != nil {
			log.Fatalln(err)
		}
		downloadUrl(r, dir)
	}
	wg.Done()
}

func main() {
	flag.Parse()
	task := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(*pool)

	tmp, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Fatalln(err)
	}
	defer os.RemoveAll(tmp)

	for i := 0; i < *pool; i++ {
		go work(task, wg, tmp)
	}

	f, err := os.Create(*file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	urls, err := fetchUrl(*site)
	if err != nil {
		log.Fatalln(err)
	}

	for _, url := range urls {
		task <- url
	}

	close(task)
	wg.Wait()

	err = makeTar(tmp, f)
	if err != nil {
		log.Fatalln(err)
	}
}
