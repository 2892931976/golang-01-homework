package main

import "fmt"

// 1 1 2 3 5 8 13 21 ......,输出100以内的斐波那契数
func main() {
	var a1, a2, sum int64
	var array = []int64{1, 1}
	a1 = 1
	a2 = 1
	for {
		if len(array) > 100 {
			fmt.Println(sum)
			break
		}
		// 也可以使用 这种特殊写法: a1,a2 = a2,a1 + a2
		sum = a1 + a2
		//fmt.Println(sum)
		array = append(array, sum)
		a1 = a2
		a2 = sum
	}
}
