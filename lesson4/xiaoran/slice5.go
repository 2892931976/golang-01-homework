package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSice("a", a)

	b := make([]int, 0, 5)
	printSice("b", b)

	c := b[:2]
	printSice("c", c)

	d := c[2:5]
	printSice("d", d)

}

func printSice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
