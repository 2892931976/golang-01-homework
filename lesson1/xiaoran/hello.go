package main

import "fmt"

func main() {
	var x int
	x = 1 //把1赋给变量x
	var y int
	y = 2        //把2赋给变量y
	swap(&x, &y) //把x和y的内存地址作为参数传给swap函数
	fmt.Println(x, y)
}

func swap(p *int, q *int) {
	//指针变量p里面存的是变量x的内存地址，指针变量q存的是y的内存地址
	var t = *p
	//p里面存的是x的内存地址，*p是取指针指向的内容，*p也就是变量x里面存的值1，1赋给t
	*p = *q
	//q里面存的y的内存地址，*q是变量y里面存的值2，把2赋值给x变量里面的值，x里面的值变成了2
	*q = t
	//变量t里面存的值是1，1赋给了y里面的值，y的值变成了1
	//最终结果，x是2，y是1
}
