/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%
9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		x int
		y int
	)
	for {
		if 1 <= len(os.Args) && len(os.Args) < 4 {
			fmt.Println("对不起，您输入的参数无效！请重新输入！\n")
			return
		}
		x, _ = strconv.Atoi(os.Args[1]) //将字符串转换成数字
		y, _ = strconv.Atoi(os.Args[3])
		switch os.Args[2] {
		case "+":
			fmt.Printf("运算结果是：%d\n", x+y)
			return
		case "-":
			fmt.Printf("运算结果是：%d\n", x-y)
			return
		case "*":
			fmt.Printf("运算结果是：%d\n", x*y)
			return
		case "/":
			fmt.Printf("运算结果是：%d\n", x/y)
			return
		case "%":
			fmt.Printf("运算结果是：%d\n", x%y)
			return
		default:
			fmt.Printf("对不起，您输入的参数无效！请重新输入！\n")
			return
		}
	}
}
