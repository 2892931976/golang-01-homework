package main

import "fmt"

func print() {
	//defer 在函数将要返回前执行
	defer func() {
		fmt.Println("defer")
	}()
	fmt.Println("hello")
}

func main() {
	print()
}
