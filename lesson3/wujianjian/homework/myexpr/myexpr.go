package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//借助os包获取命令行参数
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
	}

	switch os.Args[2] {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("myexpr: division by zero")
		} else {
			fmt.Println(a / b)
		}
	case "%":
		fmt.Println(a % b)
	default:
		fmt.Println("请输入正确的操作符!!!")
	}
}
