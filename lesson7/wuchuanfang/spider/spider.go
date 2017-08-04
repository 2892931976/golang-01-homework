package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func clearUrls(u string, urls []string) []string {
	var realurl []string
	var path2 string
	u1, err1 := url.Parse(u)
	if err1 != nil {
		log.Fatal(err1)
	}
	path1 := strings.Split(u1.Path, "/")
	path := strings.SplitAfterN(u1.Path, "/", len(path1))
	for i := 0; i < len(path1)-1; i++ {
		path2 = path2 + path[i]
	}

	for _, link := range urls {
		u2, err2 := url.Parse(link)
		if err2 != nil {
			log.Fatal(err2)
		}

		if link != "" && u2.Scheme != "" {
			realurl = append(realurl, link)
			continue
		} else if u2.Host != "" {
			link = u1.Scheme + ":" + link
		} else if u2.Path != "" && u2.Path[0] == '/' {
			link = u1.Scheme + "://" + u1.Host + link
		} else if u2.Path != "" && u2.Path[0] != '/' {
			link = u1.Scheme + "://" + u1.Host + path2 + link
		} else {
			continue
		}
		realurl = append(realurl, link)
	}

	return realurl
}

func dup(a []string) (ret []string) {
	sort.Strings(a)
	alen := len(a)
	for i := 0; i < alen; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return ret
}

func fetch(url string) ([]string, error) {
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
	var link []string
	//var ok bool
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		inlink, ok := s.Attr("src")
		if ok {
			link = append(link, inlink)
		}
	})
	link = dup(link)
	return link, nil
}

func main() {
	//url := "http://daily.zhihu.com/"
	url := "http://59.110.12.72:7070/golang-spider/img.html"
	//url := os.Args[1]
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pre url:")
	for _, u := range urls {
		fmt.Println(u)
	}

	fmt.Println("Cleard url:")
	s := clearUrls(url, urls)
	for _, u := range s {
		fmt.Println(u)
	}
}
