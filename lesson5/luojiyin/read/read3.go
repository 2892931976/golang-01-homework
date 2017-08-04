package main

import (
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("output.txt", b, 0664)
	if err != nil {
		panic(err)
	}
}
