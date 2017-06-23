package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, err := strconv.Atoi(os.Args[1])
	b, error := strconv.Atoi(os.Args[3])
	switch os.Args[2] {
	case "+":
		if err == nil && error == nil {
			fmt.Println(a + b)
		}
	case "-":
		if err == nil && error == nil {
			fmt.Println(a - b)
		}
	case "*":
		if err == nil && error == nil {
			fmt.Println(a * b)
		}
	case "/":
		if err == nil && error == nil {
			fmt.Println(a / b)
		}
	case "%":
		if err == nil && error == nil {
			fmt.Println(a % b)
		}
	}
}
