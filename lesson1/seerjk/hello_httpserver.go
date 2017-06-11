package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!!!!")
}

func main() {
	// error: no "r" end with handle
	// http.HandlerFunc("/", handler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
