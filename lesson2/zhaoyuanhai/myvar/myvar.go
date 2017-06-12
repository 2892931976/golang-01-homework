package main

import "fmt"

var x = 200

func localFunc() {
	fmt.Println(x)
	fmt.Println("local,X=", &x)
}

func main() {
	x := 1

	localFunc()    //x=200, 全局变量x值为200
	fmt.Println(x) //x=1  新声明局部变量x值初始为1
	if true {
		x := 100
		fmt.Println(x) //x=100 ,if语句内新声明局部变量x值初始为100
		fmt.Println("if,x=", &x)
	}

	localFunc()                //x=200 ，全局x值为200
	fmt.Println(x)             //x=1
	fmt.Println("main,x=", &x) //main函数声明的局部变量
}
