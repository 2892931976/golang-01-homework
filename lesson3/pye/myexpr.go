package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if args == nil || len(args) < 4 {
		_usage()
		return
	}

	num1, err := strconv.Atoi(args[1])
	if err != nil {
		return
	}

	oper := args[2]

	num2, err := strconv.Atoi(args[3])
	if err != nil {
		return
	}

	switch oper {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	}

}

var _usage = func() {
	fmt.Println("USAGE: myexpr num1 + num2, support + - * /")
	fmt.Println()
	fmt.Println("* need to be escaped")
}
