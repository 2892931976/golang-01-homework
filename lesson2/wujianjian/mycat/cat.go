package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//var s,sep string
	for i := 1; i < len(os.Args); i++ {
		printFile(os.Args[i])
	}
	//printFile("a.txt")
}

//参照myecho把cat命令的功能实现
func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Print(err)
		return
	}
	//fmt.Println(string(buf))
	fmt.Print(string(buf))
}
