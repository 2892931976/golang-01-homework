package main

import "fmt"

func main() {
	//数组定义了，默认长度是不可变的
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q[2], r[2])

	q1 := [...]int{1, 2, 3, 4}
	fmt.Println(q1)

	q2 := [...]int{4: 2, 10: -1}
	fmt.Println(q2, len(q2))
}
