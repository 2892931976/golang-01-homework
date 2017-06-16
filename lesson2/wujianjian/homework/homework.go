package main

import "fmt"

var x = 200 //全局变量

func localFunc() {
	fmt.Println(x)
}

func main() {
	x := 1 //局部变量，只影响main函数内部

	localFunc()    //调用函数，取值是全局变量，所以打印值：200
	fmt.Println(x) //直接打印局部变量：1
	if true {
		x := 100
		fmt.Println(x) //程序块的局部变量：100 ,而且影响不到上一层。
	}

	localFunc()    //因为100局部变量影响不到这一层，所以还是打印值：200
	fmt.Println(x) //直接打印局部变量：1
}
