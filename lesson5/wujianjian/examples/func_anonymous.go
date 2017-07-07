package main

import (
	"fmt"
	"strings"
)

func toupper(s string) string {
	return strings.Map(func(r rune) rune {
		return r - ('a' - 'A')
	}, s)
}

func main() {
	//匿名函数只有入参和出参，没有函数名了,类似于python的lambda表达式
	fmt.Println(toupper("hello"))
}
