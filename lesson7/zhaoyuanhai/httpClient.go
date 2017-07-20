package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func cleanUrls(u string, urls []string) []string {
	var links []string
	linkState, err := url.Parse(u)
	if err != nil {
		fmt.Println(err)
	}
	scheme := linkState.Scheme
	host := linkState.Host
	path := strings.Split(linkState.Path, "/")[1]
	for _, link := range urls {
		if strings.HasPrefix(link, "//") {
			link = scheme + ":" + link
		} else if strings.HasPrefix(link, "/") {
			link = scheme + "://" + host + link
		} else if strings.HasPrefix(link, "http") {

		} else {
			link = scheme + "://" + host + "/" + path + "/" + link
		}
		links = append(links, link)
	}
	return links
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
	//	io.Copy(os.Stdout, resp.Body)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")
		if ok {
			urls = append(urls, link)
		} else {
			fmt.Println("src not found")
		}
	})
	return urls, nil
}

func main() {
	// link := "http://daily.zhihu.com/"
	link := os.Args[1]
	urls, err := fetch(link)
	if err != nil {
		log.Fatal(err)
	}
	links := cleanUrls(link, urls)
	for _, link := range links {
		fmt.Println(link)
	}

}
