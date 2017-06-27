package main

import (
	"fmt"
)

func Reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func main() {
	word := "dfdfsdfdfasfd3e21dfds"
	fmt.Println(Reverse(word))
}
