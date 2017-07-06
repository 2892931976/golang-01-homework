package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("fmt.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(f, "hello")
	fmt.Fprintln(f, "helloln")
	s := "hello"
	n := 4
	fmt.Fprintf(f, "my string is :%s n=%d\n", s, n)
	f.Close()
}
