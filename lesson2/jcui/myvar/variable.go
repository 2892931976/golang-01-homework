package main

import "fmt"

var global string = "global"

func localfunc() *string {
	var local = "local"
	fmt.Println(local)
	return &local
}

func main() {
	// 定义全局变量
	// 每个类型都有默认的初始值-零值(安全)
	var (
		x int = 100
		y float32
		z string
		p *int
		a bool
		c *int
		d **int
	)
	// 如下是定义局部变量
	i := 0       //int
	s := "hello" //string
	m, n := 0, 1 //批量初始化
	c = &x
	d = &c

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", s)
	fmt.Printf("%v %v\n", m, n)
	fmt.Println("c=", c)
	fmt.Println("*c=", *c)
	fmt.Println("d=", d)
	fmt.Println("*d=", *d)

	ss := localfunc()
	fmt.Println(*ss)

	if true {
		x, a := 10000, 2 //批量赋值
		fmt.Println(x, a)
	}
	fmt.Println(x)
	if true {
		var a int
		x, a = 10000, 2
		fmt.Println(x, a)
	}
	fmt.Println(x)
}
