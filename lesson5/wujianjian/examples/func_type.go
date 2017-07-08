package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(m, n int) int {
	return m + n
}

func sub(m, n int) int {
	return m - n
}

func mul(m, n int) int {
	return m * n
}

func div(m, n int) int {
	return m / n
}

func main() {
	/*
		比较2个函数是否一样: 参数 + 返回值
		函数类型是：一等公民

	*/
	funcmap := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	m, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[3])

	f := funcmap[os.Args[2]]
	if f != nil {
		fmt.Println(f(m, n))
	}
}
