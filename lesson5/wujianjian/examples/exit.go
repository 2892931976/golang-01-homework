package main

import (
	"fmt"
	"os"
)

func print() {
	fmt.Println("hello")
}

func main() {
	//函数遇到return就退出函数，for遇到break就退出，其他的可以使用os.Exit(n)进行处理。
	_, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
		return
	}
	print()
}
