package main

import (
	"fmt"
	"os"
)

func main() {
	strs := os.Args[1]
	nbyte := []rune(strs)
	sz := len(nbyte)

	for i, j := 0, sz-1; i < j; i, j = i+1, j-1 {
		nbyte[i], nbyte[j] = nbyte[j], nbyte[i]
	}

	fmt.Println(string(nbyte))

	/* output:
	$go run myreverse.go ABCDEFG
	$GFEDCBA
	*/
}
