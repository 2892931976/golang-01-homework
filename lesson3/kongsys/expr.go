package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		m := convert(os.Args[1])
		n := convert(os.Args[3])
		switch os.Args[2] {
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
		fmt.Println("parameter not enough.")
	}
}

func convert(s string) int {
	temp, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, "type convert error.")
		os.Exit(1)
	}
	return temp
}
