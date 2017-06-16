package main

import "fmt"

var x = 200 //定义全局变量 x=200

func localFunc() {
	fmt.Println(x)
}

func main() {
	x := 1 //在main语段中定义局部变量x=1

	localFunc()    // 引用全局变量x,输出为200
	fmt.Println(x) //这里调用的是局部变量x，输出为1
	if true {
		x := 100       //在if语段中定义新的局部变量x=100
		fmt.Println(x) //输出调用x为100,if结束,变量失去作用
	}

	localFunc()    //调用全局变量x，输出为200
	fmt.Println(x) //调用main语段中定义局部变量x,输出为1
}
