package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello,world")
}

func main() {
	addr := ":8080"
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
