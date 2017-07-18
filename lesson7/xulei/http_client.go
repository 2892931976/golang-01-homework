package main

import (

	"net/http"

	"os"
	//"io"
	"log"
 //       "net/url"

	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func cleanUrls(u string, urls []string) ([]string, error) {
      
	
      uu, err := url.Parses(u)
      var u2 []string
      
      for  i := 0; i< len(urls); i++ {              
   
               uuu, err := url.Parses(urls[i])
               if  err != nil {

                    log.Fatal(err) 
                                  
       
               }
               if uuu.Scheme != nil {
                  contine
               }else{
                  uuu.Scheme := uu.Scheme
               }
               
               if uuu.Host != nil {
                  contine
               }else{
                  uuu.Host = uu.Host
               }
               
      return u2, nil
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
	doc.Find("img").Each(func(i int, s*goquery.Selection) {

		link, ok := s.Attr("src")
		if ok {

			urls = append(urls, link)


		}else {

			fmt.Println("src not found")

		}
	})

    return  urls, nil
}



func main() {
	 url := os.Args[1]
	 urls, err := fetch(url)
	 if err != nil {
		   log.Fatal(err)

	 }
         aa, err := cleanUrls(url, urls)
         fmt.Println(aa)
        /*
	for _, u := range urls {

		fmt.Print(u)
	}
        */
    





}
