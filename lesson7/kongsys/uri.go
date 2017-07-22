package main

import (
	"log"
	"net/url"
	"os"
)

func main() {
	s := os.Args[1]
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("scheme", u.Scheme)
	log.Println("host", u.Host)
	log.Println("path", u.Path)
	log.Println("quertString", u.RawQuery)
	log.Println("user", u.User)
	log.Println("xx", u.Fragment)
}
