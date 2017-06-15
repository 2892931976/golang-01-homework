package main

import "fmt"

func main() {
	var x int
	x = 1 //把1赋值给x
	var y int
	y = 2        //把2赋值给y
	swap(&x, &y) //把x和y的内存地址当做参数传递给参数swap
	fmt.Println(x, y)
}

func swap(p *int, q *int) {
	//指针变量p是x的内存地址，指针变量q是y的内存地址
	var t = *p
	//p里面是x的内存地址，*p是指针所指向的内容，*p也就是变量x的内容1，就是把1赋值给t
	*p = *q
	//q里面是y的内存地址，*q是指针所指向内容，*q所指向的内容是y的值，*p 指针指向的内容就是y里面的值

	*q = t
	//变量t的值是1，就是把t的值赋值给*q所指向的内容

}
