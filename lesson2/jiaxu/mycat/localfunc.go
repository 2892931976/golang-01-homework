package main

import "fmt"

var x = 200 //定义全局变量x的值为200

func localFunc() {
	fmt.Println(x) //打印结果为200
}

func main() {
	x := 1 //初始化x变量为1

	localFunc()
	fmt.Println(x) //打印结果为1
	if true {      //结果为真，将x的值赋值100，并打印出来
		x := 100
		fmt.Println(x)
	}

	localFunc()
	fmt.Println(x) //打印局部变量x:=1的值
}

//打印结果为200 1 100 1
