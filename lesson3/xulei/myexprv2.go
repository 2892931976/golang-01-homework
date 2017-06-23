package main

import (
	"fmt"
	"os"
	"strconv"
)

func myexprv2() int {
	if len(os.Args) != 4 {
		fmt.Println("please input 3 lie,example x + y")
		os.Exit(250)
	}
	var z int
	x, err1 := strconv.Atoi(os.Args[1])

	y, err3 := strconv.Atoi(os.Args[3])
	if err1 == nil && err3 == nil && len(os.Args) == 4 {

		switch os.Args[2] {

		case "+":
			z = x + y
		case "-":
			z = x - y
		case "*":
			z = x * y
		case "/":
			z = x / y
		}
	} else {
		fmt.Println("err!!,input error")
	}
	return z
}

func main() {
	fmt.Println(myexprv2())
}
