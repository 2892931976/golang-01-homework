package main

import (
	"fmt"
)

func r_slice(_slice []int, n int) []int {
	return append(_slice[n:], _slice[:n]...)
}

func main() {
	s := []int{4, 5, 6, 9, 40, 65, 74, 90, 123}
	n := 3

	if n > 0 && n < len(s) {
		fmt.Println(s)
		fmt.Println(r_slice(s, n))
	} else {
		fmt.Println("invalid length, short than slice length or greater than 0")
	}

}
