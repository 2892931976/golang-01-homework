package main

import "fmt"

func sum(args ...int) int { //用来聚拢,传入多个参数
	n := 0
	for i := 0; i < len(args); i++ {
		n += args[i]
	}
	return n
}

func main() {
	fmt.Println(sum(1, 2, 3))
	s := []int{1, 2, 3}
	fmt.Println(sum(s...)) //用来发散，列出所有值
}
