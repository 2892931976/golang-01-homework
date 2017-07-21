package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func cleanUrls(s string, urls []string) []string {
	var slice_url []string
	url_str, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range urls {
		u, err := url.Parse(v)
		if err != nil {
			log.Fatal(err)
		}
		//先判断有没协议，有协议，直接继续下次循环
		if u.Scheme != "" {
			slice_url = append(slice_url, v)
			continue
		}

		str := []byte(v)

		if string(str[0]) == "/" && string(str[1]) == "/" {
			new_url := url_str.Scheme + ":" + v
			slice_url = append(slice_url, new_url)

		} else if string(str[0]) == "/" && string(str[1]) != "/" {
			new_url := url_str.Scheme + "://" + url_str.Host + v
			slice_url = append(slice_url, new_url)
		} else {
			//这里去掉路径的最后一个斜线“/”
			new_path := strings.Split(url_str.Path, "/")
			new_path = new_path[:len(new_path)-1]
			url_str.Path = strings.Join(new_path, "/")
			new_url := url_str.Scheme + "://" + url_str.Host + url_str.Path + "/" + v
			slice_url = append(slice_url, new_url)
		}
	}
	return slice_url
}

func fetch(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
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
	url := "http://59.110.12.72:7070/golang-spider/img.html"
	//url := os.Args[1]
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	cleanurls := cleanUrls(url, urls)

	for _, u := range cleanurls {
		fmt.Println(u)
	}
}
