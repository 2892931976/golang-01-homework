package main

import "fmt"

var x = 200 //定义全局变量

func localFunc() {
	fmt.Println(x)
	fmt.Println("local,X=", &x)
}

func main() {
	x = 1 //对于使用:=定义的变量，如果新变量x与那个同名已定义变量 (这里是那个全局变量x)不在一个作用域中时，那么golang会新定义这个变量x，遮盖住全局变量x

	localFunc()    //x=200, 全局变量x值为200
	fmt.Println(x) //x=1  新声明局部变量x值初始为1
	if true {
		x = 100
		fmt.Println(x) //x=100 ,if语句内新声明局部变量x值初始为100
		fmt.Println("if,x=", &x)
	}

	localFunc()                //x=200 ，全局x值为200
	fmt.Println(x)             //x=1
	fmt.Println("main,x=", &x) //main函数声明的局部变量
}
