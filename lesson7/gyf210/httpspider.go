package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

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

func cleanUrls(u string, urls []string) ([]string, error) {
	var r []string
	s, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	scheme, host, path := s.Scheme, s.Host, s.Path
	for _, value := range urls {
		n := checkUrl(value)
		switch n {
		case 0:
			r = append(r, value)
		case 1:
			l := scheme + ":" + value
			r = append(r, l)
		case 2:
			l := scheme + "://" + host + value
			r = append(r, l)
		case 3:
			l := scheme + "://" + host + path + value
			r = append(r, l)
		}
	}
	return r, nil
}

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
		l, ok := s.Attr("src")
		if ok {
			urls = append(urls, l)
		}
	})
	return urls, nil
}

func main() {
	url := os.Args[1]
	urls, err := fetch(url)
	if err != nil {
		log.Fatalln(err)
	}
	r, err := cleanUrls(url, urls)
	if err != nil {
		log.Fatalln(err)
	}
	for _, value := range r {
		fmt.Println(value)
	}
}
