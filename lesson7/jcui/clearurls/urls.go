package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func cleanUrls(l string, urls []string) []string {
	var resurls []string
	/*
		清洗url的形式

		http://xxx.com/a.jpg
		//xx.com/a.jpgs
		/ststic/a.jpg
		a.jpg
	*/
	s_u, err := url.Parse(l)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range urls {
		u, err := url.Parse(v)
		if err != nil {
			log.Fatal(err)
		}
		//判断是否有http
		if u.Scheme != "" {
			resurls = append(resurls, v)
			continue
		}
		//判读是否有host
		if u.Host != "" {
			resurls = append(resurls, s_u.Scheme+":"+v)
			continue
		}
		//判断path
		if u.Path != "" {
			tmp_s_url := strings.SplitAfter(s_u.Path, "/")
			tmp_d_url := strings.SplitAfter(u.Path, "/")
			for i := 0; i < len(tmp_d_url)-1; i++ {
				if tmp_s_url[i] != tmp_d_url[i] && tmp_d_url[0] == "/" {
					resurls = append(resurls, s_u.Scheme+"://"+s_u.Host+strings.Join(tmp_s_url[:i], "")+strings.Join(tmp_d_url[i:], ""))
					break
				}
				if tmp_d_url[0] != "/" && tmp_s_url[i+1] != tmp_d_url[i] {
					resurls = append(resurls, s_u.Scheme+"://"+s_u.Host+strings.Join(tmp_s_url[:i+1+1], "")+strings.Join(tmp_d_url[i:], ""))
					break
				}
			}
			continue
		}

		//if u.Scheme == "" {
		//	if u.Host == "" {
		//		if u.Path == "" {
		//			tmpsurl := strings.SplitAfter(s_u.Path, "/")
		//			tmpul := strings.SplitAfter(u.Path, "/")
		//			for i := 0; i < len(tmpul)-1; i++ {
		//				if tmpul[0] != "/" {
		//					if tmpsurl[i+1] != tmpul[i] {
		//						resurls = append(resurls, s_u.Scheme+"://"+s_u.Host+strings.Join(tmpsurl[:i+1+1], "")+strings.Join(tmpul[i:], ""))
		//					}
		//				} else {
		//					if tmpsurl[i] != tmpul[i] {
		//						resurls = append(resurls, s_u.Scheme+"://"+s_u.Host+strings.Join(tmpsurl[:i], "")+strings.Join(tmpul[i:], ""))
		//					}
		//				}
		//			}
		//		}
		//	} else {
		//		resurls = append(resurls, s_u.Scheme+":"+v)
		//	}
		//} else {
		//	resurls = append(resurls, v)
		//}

	}

	return resurls
}

func fetch(l string) ([]string, error) {
	var urls []string
	resp, err := http.Get(l)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("http error code:%s", resp.Status)
		return nil, err
	}
	//io.Copy(os.Stdout, resp.Body)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := (s.Attr("src"))
		if ok {
			urls = append(urls, link)
		} else {
			fmt.Println("src not found")
		}
	})
	return cleanUrls(l, urls), err
	//return urls, err
}

func main() {
	l := "http://59.110.12.72:7070/golang-spider/img.html"
	//l := "http://daily.zhihu.com/"
	/*
		http://59.110.12.72:7070/golang-spider/img.html
	*/
	//url := os.Args[1]
	urls, err := fetch(l)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range urls {
		fmt.Println("结果:", u)
	}
}
