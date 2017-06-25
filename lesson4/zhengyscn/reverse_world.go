package main

import (
	"fmt"
	"strings"
)

func ReverseWrold(str string) {
	s := strings.Split(str, " ")
	length := len(s) - 1
	for i, _ := range s[:length/2] {
		t := s[i]
		s[i] = s[length-i]
		s[length-i] = t
	}
	fmt.Println(s)
}

func main() {
	s := "hello world nihao"
	ReverseWrold(s)
}
