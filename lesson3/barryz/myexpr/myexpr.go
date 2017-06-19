package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	start, op, end := os.Args[1], os.Args[2], os.Args[3]
	res, err := calc(start, op, end)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res)
}

func calc(start, op, end string) (ret interface{}, err error) {
	s, err := strconv.Atoi(start)
	if err != nil {
		return
	}
	e, err := strconv.Atoi(end)
	if err != nil {
		return
	}

	switch op {
	case "+":
		ret = s + e
	case "-":
		ret = s - e
	case "*":
		ret = s * e
	case "/":
		if e == 0 {
			err = fmt.Errorf("zero division error")
		} else {
			ret = s / e
		}
	default:
		err = fmt.Errorf("unknown operator")
	}

	return
}
