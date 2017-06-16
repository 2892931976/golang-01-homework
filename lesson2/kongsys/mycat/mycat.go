package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		printFile(os.Args[i])
	}
}

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(buf))
	// fmt.Println(string(buf)) // some file end is \n
}
