package main

import (
	"log"
	"net/http"
	//"io"
	//"os"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

func fetch(url string) ([]string, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	var urls []string

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")
		if ok {
			//fmt.Println(link)
			urls = append(urls, link)
			//return urls
		} else {
			fmt.Println("not found")
		}
	})
	return urls, nil
}

/*
---- https://pic1.zhimg.com/v2-58e318de6172810c1b3c7236e8e0dbb4.jpg
---- //pic4.zhimg.com/v2-40becd4a519329198ecb3807f342fd7b.jpg
---- /golang-spider/img/a.jpg
---- img/b.jpg
*/

func cleanUrl(ur string, urls []string) []string {
	var urll []string
	u, err := url.Parse(ur)
	if err != nil {
		fmt.Println("base url parse err")
	}
	base_scheme := u.Scheme
	base_host := u.Host
	base_path := strings.Split(u.Path, "/")[1]
	for _, i := range urls {
		if strings.HasPrefix(i, "http") {

		} else if strings.HasPrefix(i, "//") {
			i = "https" + i
		} else if strings.HasPrefix(i, "/") {
			i = base_scheme + "://" + base_host + "/" + i
		} else {
			i = base_scheme + "://" + base_host + "/" + base_path + "/" + i
		}
		urll = append(urll, i)
	}
	return urll
}

func main() {
	url := "http://59.110.12.72:7070/golang-spider/img.html"
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	url_list := cleanUrl(url, urls)
	fmt.Println(url_list)
}
