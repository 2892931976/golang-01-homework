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
	name := os.Args
	if len(name) < 2 {
		fmt.Println("没有找到文件")
		return
	}
	for i := 1; i < len(name); i++ {
		printFile(name[i])
	}
}
