package main

import (
	"fmt"
	"strings"
)

//数字按切片反转
func myreverse(s []int, v int) {
	fmt.Println(s, v)
	v1 := s[:v]
	v2 := s[v:]
	s = append(v2, v1...)
	fmt.Println(s)
}

//字符串按单词反转
func reverse_word(s string) {
	fmt.Println(s)
	res := strings.Fields(s)
	for x, y := 0, len(res)-1; x < y; x, y = x+1, y-1 {
		res[x], res[y] = res[y], res[x]
	}
	fmt.Println(strings.Join(res, " "))
}

func main() {
	var text = []int{2, 3, 5, 7, 11}
	v := 2
	myreverse(text, v)

	var text1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	v1 := 3
	myreverse(text1, v1)

	var word = "hello world one two 三"
	reverse_word(word)
}
