package main

import (
	"fmt"
	"os"
)

/*
指定一个切片的截断长度，旋转截断后的两个切片的位置

如切片为[2, 3, 5, 7, 11]，长度为2，则旋转后的切片为[5, 7, 11, 2, 3]]
*/

func main() {
	a := []int{2, 3, 5, 6, 7, 8}
	target := 2

	if target > len(a) {
		fmt.Println("target greater than the params's length.")
		os.Exit(1)
	}

	fmt.Println(sliceJoin(a, target))
	// output [5, 6, 7, 8, 2, 3]
	fmt.Println(sliceJoin2(a, target))
	// output [5, 6, 7, 8, 2, 3]
}

func sliceJoin(s []int, target int) []int {
	return append(s[target:], s[:target]...)
}

func sliceJoin2(s []int, target int) []int {
	sz := len(s)
	for target != 0 {
		target--
		for i := 0; i < sz-1; i++ {
			s[i], s[i+1] = s[i+1], s[i]
		}
	}

	return s
}
