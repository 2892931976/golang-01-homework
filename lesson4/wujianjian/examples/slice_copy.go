package main

import "fmt"

func main() {
	b := []int{1, 2, 3, 4}
	s := make([]int, 5)
	fmt.Printf("s:%v\n", s)

	n := copy(s, b[1:4])
	fmt.Printf("%d copied, n:%v\n", n, s)
}
