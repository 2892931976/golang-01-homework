package main

import "fmt"

func main() {
	var x int
	x = 1
	var y int
	y = 3
	var p *int
	p = &x

	*p = 2
	fmt.Println("*p=", *p)
	fmt.Println("p=", p)

	p = &y
	fmt.Println("*p=", *p)
	fmt.Println("p=", p)

	fmt.Println("x=", x)
	addpointer(&x)
	fmt.Println("x=", x)

	fmt.Println("Hello, golang!!")
}

func addpointer(r *int) {
	*r = *r + 1
}

func add(q int) {
	q = q + 1
}
