package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.args
	lenargs := len(args)

	switch {
	case lenargs < 4:
		fmt.Println("warning:参数是2个表达式")
		return
	case lenargs > 4:
		fmt.Println("参数多了")
		return
	}
	//判断输入
	a1, err1 := strconv.Atoi(args[1])
	exprtype := args[2]
	a3, err3 := strconv.Atoi(args[3])
	if err1 != nil || err3 != nil {
		fmt.Println("参数有误")
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
