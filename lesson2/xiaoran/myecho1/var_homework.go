package main

import "fmt"

var x = 200 //声明x是全局变量

func localFunc() {
	fmt.Println(x)
}

func main() {
	x := 1 //这里的x是main函数里的，局部变量

	localFunc()    //200,调用localFunc函数，里面的x引用的是全局变量，var声明的200
	fmt.Println(x) //1，这里x是main函数声明的局部变量，x的值1
	if true {
		x := 100       //这里又声明了x变量，只在if作用域有效
		fmt.Println(x) //100，这里的x是if作用域，x的值100
	}

	localFunc()    //200，调用localFunc函数，还里x还是全局变量，x的值没改，200
	fmt.Println(x) //1，这里的x是main函数的，x的值是1
}
