package main

import (
	"fmt"
	"os"
)

func reverseOne(n []rune) string {
	sz := len(n)
	for i, j := 0, sz-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}

	return string(n)
}

func reverseTwo(n []rune) string {
	sz := len(n)
	for i := 0; i < sz/2; i++ {
		mIndex := sz - i - 1
		n[i], n[mIndex] = n[mIndex], n[i]
	}

	return string(n)
}

func main() {
	strs := os.Args[1]
	nbyte := []rune(strs)

	fmt.Printf("Reverse Once: %s\n", reverseOne(nbyte))
	fmt.Printf("Reverse Twice: %s\n", reverseTwo(nbyte))

	/* output:
	$go run myreverse.go ABCDEFG
	$GFEDCBA
	*/
}
