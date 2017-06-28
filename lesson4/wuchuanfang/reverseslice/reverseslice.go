package main

import (
	"fmt"
	"strings"
)

func main() {

	var i, j int
	i = 3
	s := []int{2, 3, 5, 7, 11, 14, 16}
	s1 := s[:i]
	s2 := s[i:]
	s2 = append(s2, s1...)
	fmt.Printf("Pre slice is: %v\n", s)
	fmt.Printf("When len is: %d\n", i)
	fmt.Printf("Now slice is: %v\n", s2)

	words := "hello world nice is Golang "
	word := strings.Fields(words)
	for j = (len(word) - 1); j < len(word) && j >= 0; j-- {
		fmt.Printf("%v ", word[j])
	}

}
