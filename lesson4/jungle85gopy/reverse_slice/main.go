package main

import (
	"fmt"
	"log"
)

func revSlice(s []int, idx int) []int {
	slen := len(s)
	if idx < 0 || idx > slen {
		log.Fatal("idx out of range")
	}
	// [2, 3, 5, 7, 11]`，长度为2，则旋转后的切片为`[5, 7, 11, 2, 3]`
	rst := s[idx:]
	for j := 0; j < idx; j++ {
		rst = append(rst, s[j])
	}
	return rst
}

func main() {
	fmt.Println("reverse int slice")
	var s = []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	fmt.Println("ori:\t", s, "\nlen\t result:")
	for i := 0; i < len(s); i++ {
		fmt.Println(i, "\t", revSlice(s, i))
	}
}
