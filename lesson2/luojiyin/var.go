package main

import "fmt"

var x = 200 //在函数外面定义了一个变量， 所以x是全局变量

func localFunc() {
	fmt.Println(x)
}

func main() {
	x := 1 //在函数里面定义了一个变量，所以这个x是局部变量，作用于main（）

	localFunc()    //localFunc()是被调用的函数，是读取不到main()局部变量,只能读取到全局变量
	fmt.Println(x) //
	if true {
		x := 100       //更新局部变量的值为100 他的作用域是if或者for 这样的控制流内部
		fmt.Println(x) //if 外的局部变量x 的就恢复成 1
	}

	localFunc()    //localFunc()是被调用的函数，是读取不到main()局部变量,只能读取到全局变量
	fmt.Println(x) //x为100，是在if控制语句中，出来了就恢复成1
}
