package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	fmt.Print("guess a number 1-10:")
	var n int
	fmt.Scanf("%d", &n)

	//补全代码,如果n==x 输出正确
	//如果n>x输出过大
	//如果n<x输出过小
	if n == x {
		fmt.Println("输出正确")
	} else if n > x {
		fmt.Println("输出过大")
	} else {
		fmt.Println("输出过小")
	}
}
