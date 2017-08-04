package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	_ "reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// 1.获取链接，下载链接
func fetch(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//resp.Body必须关闭，不然会造成资源泄漏
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	//输出到屏幕，不占用内存
	//io.Copy(os.Stdin, resp.Body)

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("src")
		//link_type := reflect.TypeOf(link)
		//fmt.Println(link_type)
		urls = append(urls, link)
	})
	urls = cleanUrls(url, urls)
	return urls, nil
}

// 2.清洗url
func cleanUrls(weburl string, urls []string) []string {
	var resurls []string
	u, err := url.Parse(weburl)
	if err != nil {
		log.Fatal(err)
	}
	s := make([]string, 10)
	s = strings.Split(u.Path, "/")
	s = s[:len(s)-1]
	var path string
	for _, v := range s {
		path += v + "/"
	}
	for _, v := range urls {
		switch {
		case strings.HasPrefix(v, "http") || strings.HasPrefix(v, "https"):
			resurls = append(resurls, v)
		case strings.HasPrefix(v, "//"):
			resurls = append(resurls, u.Scheme+":"+v)
		case strings.HasPrefix(v, "/"):
			resurls = append(resurls, u.Scheme+":/"+u.Host+v)
		default:
			resurls = append(resurls, u.Scheme+":/"+u.Host+path+v)
		}

	}
	return resurls
}

// 3.下载(并行下载和串行下载)

func main() {
	url := "http://59.110.12.72:7070/golang-spider/img.html"
	//url := "http://daily.zhihu.com/"
	//url := os.Args[1]
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range urls {
		fmt.Println(u)
	}
}
