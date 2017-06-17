package main

import "fmt"

/*
变量
	1. 作用域
	2. 类型
	2. 生命周期
*/

/*
***** 局部变的优先级最大 *****
 */

// 全局变量。
// 定义int型变量x 并赋值200。
// 变量x的生命周期整个程序退出。
var x = 200

func localFunc() {
	// 打印变量x
	// 因为函数localFunc不接收任何参数，那么变量x唯一的可能就是引用全局变量x.
	// x=200
	fmt.Println(x)
}

func main() {
	// 局部变量。
	// 定义int型变量x，并赋值1。
	// 变量x的生命周期是在main函数结束。
	x := 1

	localFunc()

	// 局部变量的优先级最大.
	// 因此x=1的优先级大于x=200
	// x = 1
	fmt.Println(x)
	if true {
		// 局部变量
		// 生命周期if语句块结束
		x := 100
		// x = 100
		fmt.Println(x)
	}

	/*
		1. if语句块结束后，x=100的变量已经不存在了。
		2. x=1的优先级大于x=200
	*/
	localFunc()
	// x = 100
	fmt.Println(x)
}
