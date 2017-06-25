package main

import (
	"fmt"
)

func myslice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func mystring(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := []int{2, 3, 5, 7, 11}
	fmt.Println(a)
	myslice(a[:2])
	myslice(a[2:])
	myslice(a)
	fmt.Println(a)
	fmt.Println()

	str := "hello world"
	fmt.Println(str)
	b := []byte(str)
	mystring(b)
	fmt.Println(string(b))
}
