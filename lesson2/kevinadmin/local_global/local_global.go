package main

import (
	"fmt"
)

//定义全局变量 x 并赋值 200
var x = 200

func localFunc() {
	fmt.Println(x)
}

func main() {
	//定义main函数局部变量x 并赋值为 1
	x := 1

	//打印全局变量 x值为 200
	localFunc()
	//打印main函数局部变量x 值为1
	fmt.Println(x)
	if true {
		//定义if循环局部变量x 值为100
		x := 100
		//打印if循环局部变量x 值为 100
		fmt.Println(x)
	}

	//打印全局变量x 值为200
	localFunc()
	//打印局部变量x 值为1
	fmt.Println(x)
	//最终打印的结果为
	/*
		200
		1
		100
		200
		1
	*/
}
