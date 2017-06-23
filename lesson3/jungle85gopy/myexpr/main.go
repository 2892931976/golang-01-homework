package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var op1, op2 int
	var flag string

	if len(os.Args) != 4 {
		fmt.Printf("usage: %s  2 +|-|'*'|/ 5\n", os.Args[0])
		return
	}

	op1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("%s not int", os.Args[1])
	}

	flag = os.Args[2]
	op2, err = strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("%s not int", os.Args[3])
	}

	fmt.Printf("%d %s %d = ", op1, flag, op2)
	switch flag {
	case "+":
		fmt.Println(op1 + op2)
	case "-":
		fmt.Println(op1 - op2)
	case "*":
		fmt.Println(op1 * op2)
	case "/":
		{
			if op2 != 0 {
				fmt.Println(op1 / op2)
			} else {
				fmt.Println("\nfatal: div by zero!")
			}
		}
	default:
		fmt.Println("\nunknown operator:", flag)
		return
	}
	return
}
