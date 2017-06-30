package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//数组是值拷贝，而不是引用
	//变量赋值相当于文件，数组是文件夹
	//Go有值和引用的问题
	a1 := [3]int{1, 2, 3}
	var a2 [3]int
	//值引用
	a2 = a1
	fmt.Println(a1, a2)
	fmt.Println(&a1[0], &a2[0])
	fmt.Println(unsafe.Sizeof(a1))
	fmt.Printf("%x\n", 255)
}
