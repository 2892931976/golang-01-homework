//输出100的斐波那契数之和
package main

import (
	"fmt"
)

func main() {
	var x, y int = 1, 1
	var z int = x + y
	fmt.Printf("%d\n%d\n", x, y)
	for y < 100 {
		x, y = y, y+x
		if y < 100 {
			z += y
			fmt.Println(y)
		}
	}
	fmt.Printf("100以内斐波契数之和为：%d\n", z)
}
