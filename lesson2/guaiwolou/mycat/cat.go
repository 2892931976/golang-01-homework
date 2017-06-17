package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printFile(name string) {
	//fmt.Println(name)
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}

func main() {
	Lenth := len(os.Args)
	if Lenth < 2 {
		fmt.Println("你未加所要查看的文件名,请至少添加一个需要查看的文件名！")
	}
	for i := 1; i < len(os.Args); i++ {
		printFile(os.Args[i])
	}
}
