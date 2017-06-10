package main

import (
	"fmt"
)

func main() {
	var x int
	var y int
	x = 1
	y = 2
	swap(&x, &y)
	fmt.Println("x=", x, "y=", y)

}

// swap 参数接收int类型的指针
func swap(p *int, q *int) {
	// 申明int类型变量t
	var t int
	// 将指针类型p指向的值赋值给t
	t = *p
	// 将指针类型q指向的值赋值给指针类型p指向的值
	*p = *q
	// 将t的值赋值给指针类型q指向的值，完成交换
	*q = t

}
