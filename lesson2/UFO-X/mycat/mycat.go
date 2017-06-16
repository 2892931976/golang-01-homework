package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}

func main() {

	printFile(os.Args[1])
}
