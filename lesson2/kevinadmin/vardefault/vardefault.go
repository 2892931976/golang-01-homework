package main

import "fmt"

func main() {
	var (
		x int
		y float32
		z string
		p *int
		q bool
	)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", q)

	i := 0
	s := "hello"
	k, j := 0, 1 //批量初始化
	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(k)
	fmt.Println(j)
}
