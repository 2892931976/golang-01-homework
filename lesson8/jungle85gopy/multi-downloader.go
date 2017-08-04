package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"flag"
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

	"github.com/PuerkitoBio/goquery"
)

var (
	web   = flag.String("w", "", "url to download")
	label = flag.String("l", "img", "label to download")
	dest  = flag.String("d", "img.tar.gz", "dest file to save")
	num   = flag.Int("n", 5, "routine numbers")
)

var labelAttrMap = map[string]string{
	"img":    "src",
	"script": "src",
	"a":      "href",
}

func repairLink(uri *url.URL, url string) string {
	switch {
	case strings.HasPrefix(url, "http"):
		return url
	case strings.HasPrefix(url, "//"):
		return fmt.Sprintf("%s:%s", uri.Scheme, url)
	case strings.HasPrefix(url, "/"):
		return fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, url)
	default:
		urlPath := filepath.Dir(uri.Path) + "/" + url
		return fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, urlPath)
	}
}

func cleanLinks(oriURL string, links []string) ([]string, error) {
	var result []string
	uri, err := url.Parse(oriURL)
	if err != nil {
		return nil, err
	}
	for _, url := range links {
		result = append(result, repairLink(uri, url))
	}
	return result, nil
}

func fetchOrigin(web, label string) ([]string, error) {
	var urls []string
	resp, err := http.Get(web)
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
	doc.Find(label).Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr(labelAttrMap[label])
		if ok {
			urls = append(urls, link)
		}
	})
	return cleanLinks(web, urls)
}

// capsulate the waitgroup, concurrent pool and routine.
func downloadImgs(urls []string, dir string, conNum int) error {
	if conNum <= 0 {
		conNum = 1
	}
	var wg sync.WaitGroup
	wg.Add(conNum)
	urlChan := make(chan string)
	for i := 1; i <= conNum; i++ {
		go downWorker(urlChan, dir, &wg)
	}
	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan)
	wg.Wait()
	return nil
}

// read from channel and call downloadURL()
func downWorker(urlChan chan string, dir string, wg *sync.WaitGroup) {
	for url := range urlChan {
		err := downloadURL(url, dir)
		if err != nil {
			log.Println(err)
		}
	}
	wg.Done()
}

func downloadURL(url, dir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// close resp
	defer resp.Body.Close()
	fullName := filepath.Join(dir, path.Base(url))
	fd, err := os.Create(fullName)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = io.Copy(fd, resp.Body)
	if err != nil {
		return err
	}
	log.Println("down:", fullName)
	return nil
}

func makeTar(dir, tarName string) error {
	fd, err := os.Create(tarName)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	fmt.Println("\nstart make tar...:", dir)
	baseDir := filepath.Base(dir)

	compress := gzip.NewWriter(fd)
	tr := tar.NewWriter(compress)
	defer compress.Close()
	defer tr.Close()

	filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		relName, _ := filepath.Rel(dir, name)
		fmt.Printf(" - name:%s, p:%s\n", name, relName)
		header.Name = filepath.Join(baseDir, relName)
		tr.WriteHeader(header)
		fd, err := os.Open(name)
		if err != nil {
			return err
		}
		defer fd.Close()
		io.Copy(tr, fd)
		return nil
	})
	return nil
}

func main() {
	flag.Parse()
	// web, label, dest, num
	// "http://59.110.12.72:7070/golang-spider/img.html"
	links, err := fetchOrigin(*web, *label)
	if err != nil {
		fmt.Println(err)
	}
	for i, url := range links {
		fmt.Println(i, url)
	}

	tmpDir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)
	err = downloadImgs(links, tmpDir, *num)
	if err != nil {
		log.Panic(err)
	}
	makeTar(tmpDir, *dest)
}
