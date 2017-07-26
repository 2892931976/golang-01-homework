package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

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
		case strings.HasPrefix(link, "https") || strings.HasPrefix(link, "http"):
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
}
