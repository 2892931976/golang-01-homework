package main

import "fmt"

//package定义的变量
var x = 200

func localFunc() {
	//打印package定义的变量
	fmt.Println(x)
}

func main() {
	//main里面定义的变量
	x := 1
	//函数调用的是package里面的变量,函数没有接收变量
	localFunc()
	//打印的main函数里面的x
	fmt.Println(x)

	if true {
		x := 100
		//x重新赋值打印if里面的变量
		fmt.Println(x)
	}
	//打印package里面的变量
	localFunc()
	fmt.Println(x)
}
