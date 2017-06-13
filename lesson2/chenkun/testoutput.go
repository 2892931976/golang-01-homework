package main

import "fmt"

var x = 200

func localFunc() {
	fmt.Println(x) // 需要全局变量x
}

func main() {
	x := 1

	localFunc()    // 打印200，全局变量x
	fmt.Println(x) // 打印1，局部变量x（作用域 main函数）
	if true {
		x := 100
		fmt.Println(x) // 打印100，局部变量x（作用域 if语句块）
	}

	localFunc()    // 打印200，全局变量x
	fmt.Println(x) // 打印1，局部变量x（作用域 main函数）
}
