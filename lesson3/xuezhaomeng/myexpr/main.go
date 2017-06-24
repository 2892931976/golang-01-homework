package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 4 {
		fmt.Println("params err!eg: 2 * 3 ")
	}
	if len(os.Args) < 4 {
		fmt.Println("use err! eg: 2 * 3 ")
	}

	//var total int

	params1, err := strconv.Atoi(os.Args[1])
	params2 := os.Args[2]
	params3, err := strconv.Atoi(os.Args[3])

	if err != nil {
		fmt.Println(err)
		return

	}

	switch params2 {
	case "+":
		fmt.Println(params1 + params3)
	case "-":
		fmt.Println(params1 - params3)
	case "*":
		fmt.Println(params1 * params3)
	case "/":
		fmt.Println(params1 / params3)
	case "%":
		fmt.Println(params1 % params3)

	}

}
