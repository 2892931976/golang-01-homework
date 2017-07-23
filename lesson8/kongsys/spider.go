package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	urllib "net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func fetch(url string) ([]string, error) {
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

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")

		if ok && len(link) > 0 {
			urls = append(urls, link)
		}
	})
	return urls, nil
}
func downloadImgs(urls []string, dir string) error {
	var wg sync.WaitGroup
  wg.Add(5)
  taskch := make(chan string)
  for i := 0; i < 5; i++ {
    go work(dir, taskch, &wg)
  }
	for _, url := range urls {
    taskch <- url
  }
  close(taskch)
	wg.Wait()
	return nil
}

func work(dir string, ch chan string, wg *sync.WaitGroup) {

  for url := range ch {
    resp, err := http.Get(url)
    if err != nil {
      fmt.Println(err)
      return
    }
    defer resp.Body.Close()
    _, file := filepath.Split(url)
    fname := filepath.Join(dir, file)
    f, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0644)
    defer f.Close()
    io.Copy(f, resp.Body)
  }
  wg.Done()
}

func CleanUrl(uri *urllib.URL, link string) string {
	switch {
	case strings.HasPrefix(link, "https") || strings.HasPrefix(link, "http"):
		return link
	case strings.HasPrefix(link, "//"):
		return uri.Scheme + ":" + link
	case strings.HasPrefix(link, "/"):
		return fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, link)
	default:
		p := strings.SplitAfter(uri.Path, "/")
		path := strings.Join(p[:2], "")
		return fmt.Sprintf("%s://%s%s%s", uri.Scheme, uri.Host, path, link)
	}

}
func CleanUrls(u string, urls []string) []string {
	var ret []string
	uri, _ := urllib.Parse(u)
	for i := range urls {
		ret = append(ret, CleanUrl(uri, urls[i]))
	}
	return ret
	return urls
}

func maketar(dir string, w io.Writer) error {
  basedir := filepath.Base(dir)
  compress := gzip.NewWriter(w)
  defer compress.Close()
	tr := tar.NewWriter(compress)
	defer tr.Close()
	filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		//header.Name = name
    p, _ := filepath.Rel(dir, name)
    header.Name = filepath.Join(basedir, p)
		//fmt.Printf("name=%s, header.name=%s, info.name=%s\n", name, header.Name, info.Name())
		err = tr.WriteHeader(header)
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			f, err := os.Open(name)
			if err != nil {
				return nil
			}
			defer f.Close()
			_, err = io.Copy(tr, f)
			return err
		}
		return nil
	})
	return nil
}

func fetchImages(w io.Writer, url string) {
  urls, err := fetch(url)
  if err != nil {
    log.Fatal(err) 
  }
  urls = CleanUrls(url, urls)
	tempdir, err := ioutil.TempDir(".", "spider")
	if err != nil {
		log.Panic(err)
	}
  err = downloadImgs(urls, tempdir)
  if err != nil {
    log.Panic(err) 
  }
  maketar(tempdir, w)
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fetchImages(w, r.FormValue("u"))
}
func main() {
  http.HandleFunc("/", handleHTTP)
  http.ListenAndServe(":7071", nil)
}
