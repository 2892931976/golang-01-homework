package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	//小于4个参数,告知参数错误,退出
	if len(os.Args) < 4 {
		fmt.Println("args error , for example: 1 + 2")
		os.Exit(0)
	}
	//分别获取参数参数,并将数字转换为int类型数字,如果错误则告知错误,退出
	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[3])
	if err1 != nil && err2 != nil {
		fmt.Println("args error , for example: 1 + 2")
		os.Exit(0)
	}
	myexpr := os.Args[2]
	switch myexpr {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Unsupported operation, only '+', '-', '*', '/'")
	}
}
