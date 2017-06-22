package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("%s param oper param\n", os.Args[0])
		return
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	oper := os.Args[2]
	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	switch oper {
	case "+":
		fmt.Printf("%d %s %d = %d\n", x, oper, y, x+y)
	case "-":
		fmt.Printf("%d %s %d = %d\n", x, oper, y, x-y)
	case "*":
		fmt.Printf("%d %s %d = %d\n", x, oper, y, x*y)
	case "/":
		fmt.Printf("%d %s %d = %d\n", x, oper, y, x/y)
	default:
		fmt.Println("expr can only choose one of ['+' , '-', '*', '/']")
	}
}
