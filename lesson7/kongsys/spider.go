package main

import (
  "log"
  "net/http"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  url := "http://daily.zhihu.com"
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
      log.Println(link) 
    } else {
      log.Println("src not found") 
    }
  })
}
