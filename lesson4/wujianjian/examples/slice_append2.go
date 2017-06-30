package main

import "fmt"

func main() {
	//切片的append，一定要用原来的数组接住
	s := make([]int, 0, 1)
	//_ = append(s, 1)
	s = append(s, 1)
	fmt.Println(s)
	//_ = append(s, 2)
	s = append(s, 2, 3, 4, 5)
	fmt.Println(s)

	s1 := []int{13, 14, 15}
	s = append(s, s1...)
	fmt.Println(s)
}
