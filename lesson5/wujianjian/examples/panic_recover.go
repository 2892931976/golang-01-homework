package main

import "fmt"

func print() {
	var p *int
	fmt.Println(*p)
}

func main() {
	//recover通常和defer连用，一般用于http
	//必须放在panic之前，保证程序不挂掉
	//recover只是让当前的函数不运行下去
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	//手工触发panic
	panic("不想执行下去了")

	//panic 第一种情况:分母为0
	var n int
	fmt.Println(10 / n)

	//panic 第二种情况: null指针
	print()

	//panic 第三种情况:下标超出范围
	i := 3
	var slice [3]int
	fmt.Println(slice[i])
}
