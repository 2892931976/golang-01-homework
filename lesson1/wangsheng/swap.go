package main

import (
	"fmt"
)

func main(){
	var x int
	var y int
	x = 1
	y = 2
	swap(&x,&y)
	fmt.Println("x=",x,"y=",y)
}

// 接收int的指针
func swap(p *int,q *int){
	var t int
	//作为交换的变量

	t = *p
	//t等于 p的指针
	*p = *q
	//p等于q的指针
	*q = t
	//q等于t的指针
	
}