package main

import "fmt"

func main() {
	x, y := 1, 1
	var z int = x + y
	fmt.Println(x)
	for {
		fmt.Println(y)
		x, y = y, y+x
		//fmt.Println(x, y)
		if y > 100 {
			break
		}
		z = z + y
	}
	fmt.Printf("100以内的斐波那契数之和为 : %d\n", z)

}
