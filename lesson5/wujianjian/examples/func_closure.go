package main

import "fmt"

//所有闭包可以实现的，面向对象都可以实现
func addn(n int) func(int) int {
	return func(m int) int {
		return m + n
	}
}

func main() {
	f := addn(3)
	fmt.Println(f(2))
}
