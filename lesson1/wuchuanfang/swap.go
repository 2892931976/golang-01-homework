package main

import "fmt"

func main() {
	var x, y int
	x = 1
	y = 2
  fmt.Println("Pre: ", x, y, &x, &y)
	swap(&x, &y)
	fmt.Println("Now: ", x, y, &x, &y)
}

func swap(p *int, q *int) {
	var t int
	t = *p
	*p = *q
	*q = t
}
