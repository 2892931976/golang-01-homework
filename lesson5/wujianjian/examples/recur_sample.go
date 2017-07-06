package main

import "fmt"

/*
通项公式：a(n+1) = 2 * a(n) + n
第1项:2
求第10项:a10 = 2 * a9 + 9
*/
func a(n int) int {
	if n > 0 && n <= 1 {
		return 2 //设置初值的地方 a1 = 2
	}
	return 2*a(n-1) + (n - 1) //表达式的位置
}

func main() {
	fmt.Println(a(10))
}
