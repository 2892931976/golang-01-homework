package main

import "fmt"

func main() {
	var x int
	var y int
	x = 1
	y = 3
	var p *int
	p = &x
	fmt.Println("p=", p)    //打印指针地址
	fmt.Println("*p=", *p)  //打印指针的值
	fmt.Println("x=", x)
	*p = 2
	fmt.Println("x=", x)
	add(&x)
	fmt.Println("add - x=", x)
	add_a(x)
	fmt.Println("add_a - x=", x)
	p = &y
	fmt.Println("*p=", *p)
	fmt.Println("hello golang")
}

func add(q *int)  {
	*q = *q + 1

}
func add_a(q int)  {
	q = q + 1

}