package main

import "fmt"

var x = 200

func localFunc() {
	fmt.Println(x) //全局变量对此函数期作用故x为200
}

func main() {
	x := 1

	localFunc()    //main函数中的变量作用不作用与引用的函数中的变量
	fmt.Println(x) //main函数中设置的变量作于与main函数体内
	if true {
		x := 100
		fmt.Println(x) //输出100
	}
	localFunc()    //输出200
	fmt.Println(x) //输出1 main函数内部的变量没有变化

}
