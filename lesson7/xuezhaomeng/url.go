package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CleanUrl(uri *url.URL, link string) string {
	switch {
	case strings.HasPrefix(link, "https") || strings.HasPrefix(link, "http"):
		return link
	case strings.HasPrefix(link, "//"):
		return uri.Scheme + ":" + link
	case strings.HasPrefix(link, "/"):
		return fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, link)
	default:
		p := strings.SplitAfter(uri.Path, "/")
		path := strings.Join(p[:2], "") //一般情况是这样 ,/static/img/logo.png
		return fmt.Sprintf("%s://%s%s%s", uri.Scheme, uri.Host, path, link)
	}
}

func GetImgUrl(u string) ([]string, error) {
	var urls []string
	//获取URL的response
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	//io.Copy(os.Stdout,resp.Body)

	//使用goquery获取网页中的图片链接
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")
		if ok {
			urls = append(urls, link)
		} else {
			fmt.Println("src not  found")
		}
	})
	return urls, nil
}

func main() {
	access_url := "http://59.110.12.72:7070/golang-spider/img.html"
	urls, err := GetImgUrl(access_url)
	if err != nil {
		log.Fatal(err)
	}
	access_url_info, err := url.Parse(access_url)
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range urls {
		fmt.Println(CleanUrl(access_url_info, link))
	}

}
