package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	switch {
	case len(os.Args) < 4:
		fmt.Println("args must have two ！")
		return
	case len(os.Args) > 4:
		fmt.Println("too many agrs！")
		return
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return
	}

	switch os.Args[2] {
	case "+":
		fmt.Printf("result：%d\n", x+y)
	case "-":
		fmt.Printf("result：%d\n", x-y)
	case "*":
		fmt.Printf("result：%d\n", x*y)
	case "/":
		if y == 0 {
			fmt.Printf(" zimu can not zero")
			return
		} else {
			fmt.Printf("result：%d\n", x/y)
		}
	default:
		fmt.Println("I only kown - + * /")
	}
}
