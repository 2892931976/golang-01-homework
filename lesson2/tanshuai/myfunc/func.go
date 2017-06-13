package main

import "fmt"

//全局变量
var x = 200

func main() {
	//局部变量
	x := 1

	//localFunc调用var定义的全局变量，输出结果200；
	localFunc()
	//fmt这里打印的是函数体内的变量，输出结果1;
	fmt.Println(x)

	if true {
		//(:=)更改局部变量，只对这个判断语句内生效，如果使用(=)则对整个函数生效；
		x := 100
		//输出100；
		fmt.Println(x)
	}
	//跟上面执行原理一样，调用的是全局变量，输出200;
	localFunc()
	//打印局部变量
	fmt.Println(x)
}

//外部定义函数,这个函数打印全局变量
func localFunc() {
	fmt.Println(x)
}
