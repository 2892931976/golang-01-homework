package main

import "fmt"

func main() {
	//空切片,默认值是：nil
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	a := [...]int{1, 2, 3}
	s1 := a[:0]
	fmt.Println(s1 == nil)
}
