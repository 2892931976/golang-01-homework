package main

import (
	"fmt"
	"os" //所有跟系统打交道的，包括文件操作，环境变量，命令行参数等
)

func main() {
	var s, sep string
	//Args代表一个字符串数组,len能返回数组的长度
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
}
