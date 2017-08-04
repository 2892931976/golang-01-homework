package main

import "fmt"

func main() {
	s := make([]int, 0, 1)
	fmt.Println(len(s), cap(s))
	s = append(s, 1)
	fmt.Println(s)
	s = append(s, 2, 3, 4)
	fmt.Println(s)

	s1 := []int{13, 14, 15}
	s = append(s, s1...)
	fmt.Println(s)

}
