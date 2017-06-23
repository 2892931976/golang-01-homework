package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	n2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}

	switch os.Args[2] {
	case "+":
		fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d\n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d\n", n1, n2, n1*n2)
	case "/":
		if n2 == 0 {
			fmt.Println("Error,The denominator cannot be zero")
			return
		}
		fmt.Printf("%d / %d = %v\n", n1, n2, float32(n1)/float32(n2))

	}

}
