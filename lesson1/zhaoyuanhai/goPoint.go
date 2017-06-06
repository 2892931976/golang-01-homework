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
	fmt.Println("x=",x, "y=",y) 
}

func swap(p *int, q*int) {
	var t int
	t = *p  //将指针p的值赋值给变量t
	*p = *q  //将指针q的值赋值给指针p
	*q = t  //将变量t的值即原指针p的值赋指针q
	//相当于当指针p和指针q对应的真实值做交换
}
