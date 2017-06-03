package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello world")
}

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello jcui")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}

