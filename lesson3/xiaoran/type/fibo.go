package main

import "fmt"

func main() {
	var x int
	var y int
	var sum int
	x = 1
	y = 1
	fmt.Println(x)
	fmt.Println(y)
	sum = x + y

	for {
		x, y = y, x+y
		if y > 100 {
			break
		}
		sum = sum + y
		fmt.Println(y, sum)
	}
}
