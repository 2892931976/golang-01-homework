package main

import "fmt"

var x = 200 //定义全局变量 x=200

func localFunc() {
	fmt.Println(x)
}

func main() {
	x := 1 //定义局部变量x=1

	localFunc()    // 输出为200
	fmt.Println(x) //这里调用的是局部变量，输出为1
	if true {
		x := 100       //定义x=100
		fmt.Println(x) //输出为100
	}

	localFunc()    //调用全局变量x，输出为200
	fmt.Println(x) //仍然是开始的局部变量x=1,输出为1
}
