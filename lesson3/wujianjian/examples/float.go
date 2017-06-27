package main

import "fmt"

func main() {
	var a, b int
	a = 2
	b = 1
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a = a + 3 // a += 3
	if a >= b {

	}

	s := "hello"
	s += "world"

	if (a > b && b > 3) || b > 10 {

	}

	if a == b || a != b {

	}

	fmt.Println(3 / 2)
	fmt.Println(3.0 / 2)

	var c int
	c = a << 1
	c = a * 2
	fmt.Println(c)
}
