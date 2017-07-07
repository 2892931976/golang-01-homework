package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func read(f *os.File) (string, error) {
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func main() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	var content string
	var retries uint = 3
	var i uint
	for i = 1; i <= retries; i++ {
		content, err = read(f)
		if err == nil {
			break
		}
		time.Sleep(time.Second << i)
	}
	fmt.Println(content)
}
