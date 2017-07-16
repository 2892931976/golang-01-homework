package main

import (
	"errors"
	"fmt"
	"net/http"
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

func cleanUrls(url string, urls []string) []string {
	var cleanUrls []string
	for _, v := range urls {
		//fmt.Println(k, v)
		if strings.Contains(v, "//") {
			if strings.Index(v, "http") == 0 {
				//fmt.Println("ok")
				cleanUrls = append(cleanUrls, v)
			} else {
				//fmt.Println("need add http")
				cleanUrls = append(cleanUrls, "http:"+v)
			}

		} else {
			temp := strings.SplitAfter(url, "/")
			temp1 := strings.SplitAfter(v, "/")

			if strings.Index(v, "/") == 0 {
				tempUrl := strings.Join(temp[:3], "") + strings.Join(temp1[1:], "")
				//fmt.Println(tempUrl)
				cleanUrls = append(cleanUrls, tempUrl)
			} else {
				l1 := len(temp)
				//l2 := len(temp1)
				tempUrl := strings.Join(temp[:l1-1], "") + v
				//fmt.Println(tempUrl)
				cleanUrls = append(cleanUrls, tempUrl)
			}

			/*if  len(temp1) > 2  && temp[3] == temp1[2] {
				fmt.Println(strings.Join(temp[0:2], "/") + strings.Join(temp1[2:], "/"))
			} else {*/

		}

		/*fmt.Printf("%q\n", strings.Split(v, "/"))
		temp := strings.Split(v, "/")
		fmt.Println(temp[2])

		fmt.Printf("%q\n", strings.Split(url, "/"))
		temp1 := strings.Split(url, "/")
		fmt.Println(temp1[3])*/

	}
	return cleanUrls
}

func main() {
	url := "http://59.110.12.72:7070/golang-spider/img.html"
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
	/*resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	//io.Copy(os.Stdout, resp.Body)

	/*doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")
		if ok {
			fmt.Println(link)
		} else {
			fmt.Println("src not found")
		}
	})*/
}
