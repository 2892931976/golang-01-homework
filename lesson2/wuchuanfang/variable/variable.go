package main

import "fmt"

//定义一个全局变量 i 并赋值 200，作用域为全局
var i = 200

//调用并打印全局变量 i 的值 200
func local() {
	fmt.Println(i)
}

func main() {
	//新定义一个局部变量 i 并赋值 1，作用域main函数
	i := 1
	//调用并打印全局变量 i 的值 200
	local()
	//打印局部变量 i 的值 1
	fmt.Println(i)
	if true {
		//新定义一个局部变量 i 并赋值 100，作用域if函数
		i := 100
		//打印if函数作用域的局部变量 i 的值 100
		fmt.Println(i)
	}
	//调用并打印全局变量 i 的值 200
	local()
	//调用并打印全局变量 i 的值 200
	local()
	//打印main函数作用域的局部变量 i 的值 1
	fmt.Println(i)
}
