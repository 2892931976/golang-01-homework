package main

import (
	"fmt"
	"strings"
)

type IntSlice []int

func main() {
	s := []int{1, 3, 5, 7, 9}
	fmt.Println(reverse_sliceside(s, 3))

	content := "a1 a2 a3 a4 a5"
	fmt.Println(reverse_words(content))
}

// 旋转切片两侧
func reverse_sliceside(c IntSlice, n int) IntSlice {
	s := append(c[n:], c[:n]...)
	return s
}

// 旋转单词
func reverse_words(content string) []string {
	slicea := strings.Fields(content)
	lenslicea := len(slicea)
	for i := lenslicea - 1; i >= 0; i-- {
		slicea = append(slicea, slicea[i])
	}
	return slicea[lenslicea:]
}
