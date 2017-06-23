package main

import (
	"fmt"
	"os"
)

var result string

func main() {
	str := os.Args[1]
	var array = []byte(str)
	fmt.Printf("翻转前: %s\n", str)
	for i := len(array) - 1; i >= 0; i-- {
		result = result + string(array[i])
	}
	fmt.Printf("翻转后: %s\n", result)
}
