package main

import "fmt"

func reverse(temp []int, num int) (s []int) {
	if num > len(temp)-1 {
		fmt.Println("number is too big,   num need small than then length of strings")
		return
	}
	if num < 0 {
		fmt.Println("This is too small, num is 0 or more")
		return
	}
	a1 := temp[:num]
	s = temp[num:]
	//s = a1 + a2
	for _, word := range a1 {
		s = append(s, word)
	}
	fmt.Println(s)
	return s
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//reverse(s, -1)
	reverse(s, 0)
	reverse(s, 8)
	reverse(s, 3)
	//reverse(s, 9)
}
