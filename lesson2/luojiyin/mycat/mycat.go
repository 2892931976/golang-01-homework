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
    if len(os.Args) != 2 {
        fmt.Println("need  file name")
        return
    }
    file := os.Args[1]
    printFile(file)
}
