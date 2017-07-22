package main

import (
	"io"
	"log"
	"net/http"
	"os"
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
	io.Copy(os.Stdout, resp.Body)
}
