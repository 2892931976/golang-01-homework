package main

import (
	"fmt"
	"os"
	"strconv"
)

func reverse(s []int) []int {
	var result []int

	for i := 1; len(s)-i >= 0; i++ {
		result = append(result, s[len(s)-i])
	}
	return result
}

func main() {
	s := []int{2, 3, 5, 7, 11}
	fmt.Printf("翻转前:%v\n", s)
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	s1 := s[:n]
	s1 = reverse(s1)

	s2 := s[n:]
	s2 = reverse(s2)

	s1 = append(s1, s2...)
	s1 = reverse(s1)
	fmt.Printf("翻转后:%v\n", s1)
}
