package main

import "fmt"

func main() {
	//数组的中括号中要么有值，要么使用3个点号
	//而切片的中括号中的NULL的
	//数组中元素的内存地址是连续的。
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	//fmt.Println(a[len(a)])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}
