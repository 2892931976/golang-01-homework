package main

import "fmt"

var x = 200 //定义全局变量

func localFunc() {
	fmt.Println(x) //调用全局变量
}

func main() {
	x := 1 //定义main 局部变量

	localFunc()    //调用全局变量
	fmt.Println(x) //调用main 的局部变量
	if true {
		x := 100       //定义if 局部变量
		fmt.Println(x) // 调用if 局部变量
	}

	localFunc()    //调用全局变量
	fmt.Println(x) //调用main 局部变量
}
