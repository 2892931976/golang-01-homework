package main

import (
	"net/http"
	"os"
	//"io"
	"fmt"
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	//"reflect"
)

func cleanUrls(u string, urls []string) ([]*url.URL, error) {
	oldurl, err := url.Parse(u)
	if err != nil {
		log.Fatal(err)
	}
	var xinurl []*url.URL
	for i := 0; i < len(urls); i++ {
		newurl, err := url.Parse(urls[i])
		if err != nil {
			log.Fatal(err)
		}
		// if newurl.Scheme != "" {
		// 	//fmt.Println(newurl.Scheme)
		// 	continue
		// } else {
		// 	newurl.Scheme = oldurl.Scheme
		// 	//fmt.Println(newurl.Scheme)
		// }
		// if newurl.Host != "" {
		//fmt.Println(newurl.Host)
		// 	continue
		// } else {
		// 	newurl.Host = oldurl.Host
		// 	//fmt.Println(newurl.Host)
		// }
		// if newurl.Path != "" {
		// 	//fmt.Println(newurl.Path)
		// 	continue
		// } else {
		// 	newurl.Path = oldurl.Path
		// 	//fmt.Println(newurl.Path)
		//
		// }
		if newurl.Scheme == "" {

			newurl.Scheme = oldurl.Scheme
			//fmt.Println(newurl.Scheme)
		}
		if newurl.Host == "" {

			newurl.Host = oldurl.Host
			//fmt.Println(newurl.Host)
		}
		if newurl.Path == "" {

			newurl.Path = oldurl.Path
			//fmt.Println(newurl.Path)

		}
		xinurl = append(xinurl, newurl)
	}
	return xinurl, nil

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
	url := os.Args[1]
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)

	}
	aa, err := cleanUrls(url, urls)
	if err != nil {

		log.Fatal(err)
	}
	for _, newurl := range aa {

		fmt.Printf("newurl is %v\n", newurl)
	}

}
