package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	f.Read(buf)

	fmt.Printf("###%s###\n", string(buf))
	f.Close()

}
