package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	lenargs := len(args)

	// 参数个数只能是4个
	switch {
	case lenargs < 4:
		fmt.Println("Warning:参数是2个数的算术表达式（例如: 2 + 3）")
		return
	case lenargs > 4:
		fmt.Println("Warning：参数太多了...")
		return
	}

	// 判断输入，确保算术符号两边的是数字
	a1, err1 := strconv.Atoi(args[1])
	exprtype := args[2] // 算术符号
	a3, err3 := strconv.Atoi(args[3])
	if err1 != nil || err3 != nil {
		fmt.Println("Warning：参数需要是数字！")
		return
	}

	switch exprtype {
	case "+":
		fmt.Println(a1 + a3)
	case "-":
		fmt.Println(a1 - a3)
	case "*":
		fmt.Println(a1 * a3)
	case "/":
		fmt.Println(a1 / a3)
	default:
		fmt.Println("Warning：算数符号错了！")
	}
}
