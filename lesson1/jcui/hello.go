package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	var y int
	x = 1
	y = 3
	var p *int
	p = &x
	fmt.Println("p=", p)   //打印指针地址
	fmt.Println("*p=", *p) //打印指针的值
	fmt.Println("x=", x)
	*p = 2
	fmt.Println("x=", x)
	add(&x)
	fmt.Println("add - x=", x)
	add_a(x)
	fmt.Println("add_a - x=", x) //指针作改变,但值不变
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)
	p = &y
	fmt.Println("*p=", *p)
	fmt.Println("hello golang")
	//测试打印当前时间,结果发现格式化系统的layout必须指定如下格式: 2006-01-02 15:04:05.999999999 -0700 MST
	fmt.Println(time.Now().Format("2006-01-02 15:04:05 MST"))
}

func add(q *int) {
	//值+1 当前*q = 2 ,+1之后=3
	*q = *q + 1

}
func add_a(q int) {
	fmt.Println("qqqqq = ", q)
	q = q + 1
	fmt.Println("qqqqq+1 = ", q)

}
