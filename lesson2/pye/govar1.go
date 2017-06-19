package main

import "fmt"

// 声明变量x，值为200
var x = 200

func localFunc() {
	fmt.Println(x)
	fmt.Println(&x)
}

func main() {
	// 在main()内声明变量x,此时的x的内存地址和全局不一样
	// 此时函数localfunc()输出200，调用的是全局变量x
	// fmt.Println(x)调用的是main()中的x，输出1
	x := 1
	localFunc()
	fmt.Println(x)
	fmt.Println(&x)

	if true {
		//在if函数内声明变量x，内存地址和之前也不一样，此时输出100
		x := 100
		fmt.Println(x)
		fmt.Println(&x)
	}

	//函数localfunc始终是调用全局变量x，此时输出200
	localFunc()

	//输出1，在if函数中的x = :100只能在if函数中生效。此时调用的是main()中的变量x
	fmt.Println(x)
	fmt.Println(&x)

}
