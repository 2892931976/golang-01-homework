package main

import "fmt"

// 内存框图分析下
func main() {
	var x int
	x = 1
	var y int
	y = 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(p *int, q *int) {
	var t = *p
	*p = *q
	*q = t
}
