package main

import (
	"fmt"
	"strings"
)

func reverse_varbs(str string) string {
	words := strings.Fields(str)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverse_varbs("golang shi ge  niu x de yu yan"))
}
