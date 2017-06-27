package main

import (
	"fmt"
	"os"
)

func toupper(s string) string {
	var f string
	for i := 0; i < len(s); i++ {
		f += string(byte(s[i] - 32))

	}
	return f

}

func main() {
	fmt.Println(toupper(os.Args[1]))

}
