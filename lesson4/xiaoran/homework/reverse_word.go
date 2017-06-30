package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	var result string
	var array = []byte(s)
	for i := 1; len(array)-i >= 0; i++ {
		result = result + string(array[len(array)-i])
	}

	return result
}

func main() {
	s := "hello word"
	fmt.Printf("翻转前: %s\n", s)
	var new_result string

	split_s := strings.Split(s, " ")
	for _, v := range split_s {
		new_result = new_result + " " + reverse(v)
	}

	new_result = reverse(new_result)
	fmt.Printf("翻转后: %s\n", new_result)
}
