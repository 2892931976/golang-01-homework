package main

import (
	"fmt"
)

var x = 200 //声明并初始化变量x 类型 int

func localFunc() {
	fmt.Println(x)
}

func main() {
	x := 1      //重新声明变量 x 影子变量
	localFunc() //函数内 变量x 取的是全局变量，并没有传参或者做指针引用；任何地方调用此函数 打印结果都是200
	if true {
		x := 100       //重新声明变量 x  影子变量
		fmt.Println(x) //if作用域内，x为100
	}

	localFunc()
	fmt.Println(x) //打印的是此函数内同作用域的变量值 1

	// 打印结果：200、 100、 200、 1
}
