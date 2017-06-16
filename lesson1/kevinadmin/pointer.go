package main

import "fmt"

func main() {
	//声明 变量 x 类型为 int 为整型
	var x int
	//将整型数值 1 赋给 变量 x
	x = 1
	//声明 变量 y 类型为 int 为整型
	var y int
	//将整型数值 2 赋给 变量 y
	y = 2

	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(p *int, q *int) {
	//swap接收的值，p的值为1 ， q的值为2
	//声明 变量 t， 并将p所指向的地址赋值给 t，即t的值为 1
	var t = *p
	//p指向的地址值为1，q指向的地址值为2，这里把q指向的地址值为2，赋值给p指向的地址的值，这时p指向的地址值为2
	*p = *q
	//上面经过赋值，t的值为1，这里把t的值1赋给q指向的地址，即q指向的地址的值为1
	*q = t
	//经过swap后，p的值为2 q的值为1 fmp.Println(x, y)时， x的值为2 y的值为1
}