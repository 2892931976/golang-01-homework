package main

import (
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
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
		if ok {
			//fmt.Println(link)
			urls = append(urls, link)
		} else {
			fmt.Println("src  not found")
		}
	})
	return urls, nil
}

func cleanUrls(u string, urls []string) []string {
	var cleanUrls []string
	linkState, err := url.Parse(u)
	if err != nil {
		fmt.Println(err)
	}
	scheme := linkState.Scheme
	host := linkState.Host
	for _, link := range urls {
		switch {
		case strings.HasPrefix(link, "http") || strings.HasPrefix(link, "https"):
			cleanUrls = append(cleanUrls, link)
		case strings.HasPrefix(link, "//"):
			link = scheme + ":" + link
			cleanUrls = append(cleanUrls, link)
		case strings.HasPrefix(link, "/"):
			link = fmt.Sprintf("%s://%s%s", scheme, host, link)
			cleanUrls = append(cleanUrls, link)
		}
	}
	return cleanUrls
}

func downloadimg(dir, target string) error {
	log.Print(target)
	uri, err := url.Parse(target)
	if err != nil {
		return err
	}

	resp, err := http.Get(target)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	name := path.Base(uri.Path)
	fullpath := filepath.Join(dir, name)
	log.Print(fullpath)
	f, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer f.Close()

	io.Copy(f, resp.Body)
	return nil
}

func downloadimgs(dir string, urls []string) error {
	for _, u := range urls {
		if err := downloadimg(dir, u); err != nil {
			log.Print(err)
		}
	}
	return nil
}

func maketar(dir string, w io.Writer) error {
	basedir := filepath.Base(dir)
	tw := tar.NewWriter(w)
	defer tw.Close()

	return filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		p, _ := filepath.Rel(dir, name)
		header.Name = filepath.Join(basedir, p)
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			f, err := os.Open(name)
			defer f.Close()
			_, err = io.Copy(tw, f)
			return err
		}
		return nil
	})
}

func main() {
	url := "http://www.cikers.com"
	if strings.Index(url, "http") == -1 {
		fmt.Println("Please make sure   http or https")
		return
	}
	urls, err := fetch(url)
	if err != nil {
		fmt.Println(err)
	}
	cUrls := cleanUrls(url, urls)
	for k, v := range cUrls {
		fmt.Println(k, v)
	}
	dir, err := ioutil.TempDir("", "img")
	if err != nil {
		log.Panic(err)
	}

	err = downloadimgs(dir, cUrls)
	if err != nil {
		log.Panic(err)
	}
	var fi
	err = maketar(dir, w)
	if err != nil {
		log.Panic(err)
	}
}
