package main

import (
	"fmt"
	"os"
)

func main() {
	m := 10
	n := 3
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "+":
			fmt.Println(m + n)
		case "-":
			fmt.Println(m - n)
		case "*":
			fmt.Println(m * n)
		case "/":
			fmt.Println(m / n)

		default:
			fmt.Println("unkown operator")
		}
	} else {
		fmt.Println("need operator")
	}
}
