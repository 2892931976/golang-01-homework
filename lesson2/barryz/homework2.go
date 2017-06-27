package main

import "fmt"

// package级别的全局变量x
var x = 200

func localFunc() {
	// 引用package级别的变量x
	fmt.Println(x)
}

func main() {
	// 作用域为当前main函数
	x := 1

	// output --> 200
	localFunc()
	// output --> 1, 引用当前作用域的变量x；
	fmt.Println(x)
	if true {
		x := 100
		// output --> 100, 重新声明变量x并赋值， 作用域为当前if语句块
		fmt.Println(x)
	}

	// output --> 200
	localFunc()
	// output --> 1 main函数内部的x未发生修改
	fmt.Println(x)
}
