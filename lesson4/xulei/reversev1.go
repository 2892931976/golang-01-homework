package main

import (
	"fmt"
	"strings"
)

func reverse(s []int) []int {
	//for i := 0; i < len(s)/2; i++ {
	// fmt.Println(len(s)/2)
	// s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	//}
	for x, y := 0, len(s)-1; x < y; x, y = x+1, y-1 {
		s[x], s[y] = s[y], s[x]
	}
	return s
}
func reversev1(str []string) []string {
	//word := strings.Fields(str)
	//for a, b := 0, len(word)-1; a < b; a, b = a + 1, b - 1 {
	//	word[a], word[b] = word[b], word[a]
	//}
	//return word
	//word := strings.Fields(str)
	for a, b := 0, len(str)-1; a < b; a, b = a+1, b-1 {
		str[a], str[b] = str[b], str[a]
	}
	fmt.Println(str)
	return str
}
func reversev2(str string) []string {
	word := strings.Fields(str)
	for a, b := 0, len(word)-1; a < b; a, b = a+1, b-1 {
		word[a], word[b] = word[b], word[a]
	}
	return word
	//fmt.Println(word)

	//word := strings.Fields(str)
	//for a, b := 0, len(str)-1; a < b; a, b = a + 1, b - 1 {
	//	str[a], str[b] = str[b], str[a]
	//}
	//fmt.Println(str)
	//return str
}
func main() {

	s1 := []int{2, 3, 5, 7, 11, 22, 33, 44}
	//fmt.Println(len(s1[:3]))
	reverse(s1[:3])
	fmt.Println(s1)
	reverse(s1[3:])
	fmt.Println(s1)
	reverse(s1)
	fmt.Println(s1)
	//s := []string{"hello", "world", "abc"}
	//fmt.Println(reversev1(s))
	s2 := "hello world abc"
	s3 := reversev2(s2)
	//fmt.Println(reversev2(s2))
	fmt.Println(s3)
}
