package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		var Args []string
		字符串数据
	*/
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
