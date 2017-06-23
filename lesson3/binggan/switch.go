package main

import (
	"fmt"
	"os"
)

func main() {
	m := 10
	n := 3
	switch os.Args[1] {
	case "+":
		fmt.Println(m + n)
	}
}
