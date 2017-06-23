//
// myexpr.go
// Copyright (C) 2017 root <root@localhost.localdomain>
//
// Distributed under terms of the MIT license.
//
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	op := os.Args[2]
	num1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	num2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return
	}
	//fmt.Println(num1)
	//fmt.Println(op)
	//fmt.Println(num2)
	switch op {
	case "-":
		fmt.Printf("%v - %v = %v\n", num1, num2, num1-num2)
	case "+":
		fmt.Printf("%v + %v =  %v\n", num1, num2, num1+num2)
	case "*":
		fmt.Printf("%v * %v = %v\n", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("zero cannot as divisor")
			return
		} else {
			fmt.Printf("%v / %v = %v\n", num1, num2, num1/num2)
		}
	default:
		fmt.Println("I only kown - + * /")
	}

}
