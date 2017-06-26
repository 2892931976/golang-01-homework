package main

import (
	"fmt"
)

//数字按切片反转
func myreverse(s []int, v int) {
	v1 := s[:v]
	v2 := s[v:]
	s = append(v2, v1...)
	fmt.Println(s)
}

//字符串按单词反转
func reverse_word(s string) {
	fmt.Println(s)
}

func main() {
	var text = []int{2, 3, 5, 7, 11}
	v := 2
	myreverse(text, v)

	var word = "hello world"
	reverse_word(word)
}
