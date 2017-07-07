package main

import "fmt"

func main() {
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2} //不初始化全部的，只初始化有限个

	fmt.Println(r[2])
	fmt.Println(q)

	q1 := [...]int{1, 2, 3, 4} //前面3个点，表示后面写了几个元素，数组长度就是几
	fmt.Println(q1)

	q2 := [...]int{4: 2, 10: -1} //下标是4的元素初始化为2，下标10的元素初始化为-1，其余默认值
	fmt.Println(q2)

}
