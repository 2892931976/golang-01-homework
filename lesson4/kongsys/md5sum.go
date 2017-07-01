package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			f, err := os.Open(os.Args[i])
			if err != nil {
				log.Fatal(err)
			}
			check(f)
			fmt.Printf("%s\n", os.Args[i])
			defer f.Close()
		}
	} else {
		check(os.Stdin)
		fmt.Println("stdin")
	}


}

func check(i io.Reader) {
	h := md5.New()
	if _, err := io.Copy(h, i); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x ", h.Sum(nil))
}
