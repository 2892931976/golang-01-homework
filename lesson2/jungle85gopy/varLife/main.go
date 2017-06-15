package main

import "fmt"

var x = 200

func localFunc() {
	fmt.Println(x)
}

func main() {
	x := 1

	localFunc()
	// output: 200; localFunc中的x是全局变量
	fmt.Println(x)
	// output: 1;	x是main内部的变量
	if true {
		x := 100
		fmt.Println(x)
		// output: 100; x是if语句的局部变量
	}

	localFunc()
	// output: 200; localFunc中的x是全局变量
	fmt.Println(x)
	// output: 1;	x是main内部的变量
}
