package main

import "fmt"

const (
	PI = 3.1415926
	E  = 2.0
	G  = 9.8
)

const (
	A = 10 * iota
	B
	C
)

func main() {
	var n float32
	n = PI
	fmt.Println(n)
	fmt.Println(A, B, C)
}
