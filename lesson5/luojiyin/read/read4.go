package main

import (
	"io"
	"log"
	"os"
)

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	w, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	n, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}
	w.Sync()
	log.Printf("Copied %v bytes\n", n)
}
