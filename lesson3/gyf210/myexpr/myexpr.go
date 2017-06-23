package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: myexpr 1 + 2")
		return
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch os.Args[2] {
	case "+":
		fmt.Printf("%v + %v = %v\n", x, y, x+y)
	case "-":
		fmt.Printf("%v - %v = %v\n", x, y, x-y)
	case "*":
		fmt.Printf("%v * %v = %v\n", x, y, x*y)
	case "/":
		fmt.Printf("%v / %v = %v\n", x, y, x/y)
	default:
		fmt.Println("Usage: myexpr 1 + 2")
	}
}
