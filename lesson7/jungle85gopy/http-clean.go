package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func repairLink(uri *url.URL, url string) string {
	if yes := strings.HasPrefix(url, "http"); yes {
		return url
	} else if yes = strings.HasPrefix(url, "//"); yes {
		return uri.Scheme + ":" + url
	} else if yes = strings.HasPrefix(url, "/"); yes {
		return uri.Scheme + "://" + uri.Host + url
	}
	pathS := strings.Split(uri.Path, "/")
	path := strings.Join(pathS[0:len(pathS)-1], "/") + "/" + url
	return uri.Scheme + "://" + uri.Host + path
}

func clean(oriURL string, imgLinks []string) ([]string, error) {
	var result []string
	uri, err := url.Parse(oriURL)
	if err != nil {
		return nil, err
	}
	for _, url := range imgLinks {
		result = append(result, repairLink(uri, url))
	}
	return result, nil
}

func fetch(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// close resp
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
			urls = append(urls, link)
		}
	})
	return urls, nil
}

func main() {
	url := os.Args[1]
	// url := "http://59.110.12.72:7070/golang-spider/img.html"
	imgLinks, err := fetch(url)
	if err != nil {
		fmt.Println(err)
	}
	/*
		<img src="https://pic1.zhimg.com/v2-58e318de6172810c1b3c7236e8e0dbb4.jpg" />
		<img src="//pic4.zhimg.com/v2-40becd4a519329198ecb3807f342fd7b.jpg" />
		<img src="/golang-spider/img/a.jpg" />
		<img src="img/b.jpg" />
	*/
	result, err := clean(url, imgLinks)
	if err != nil {
		fmt.Println(err)
	}
	for i, url := range result {
		fmt.Println(i, url)
	}
	/*
		0 https://pic1.zhimg.com/v2-58e318de6172810c1b3c7236e8e0dbb4.jpg
		1 http://pic4.zhimg.com/v2-40becd4a519329198ecb3807f342fd7b.jpg
		2 http://59.110.12.72:7070/golang-spider/img/a.jpg
		3 http://59.110.12.72:7070/golang-spider/img/b.jpg
	*/
}
