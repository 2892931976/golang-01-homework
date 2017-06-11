package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "你好吗")
	fmt.Println(time.Now().Format("2007-01-02 15:04:05 MST"))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
