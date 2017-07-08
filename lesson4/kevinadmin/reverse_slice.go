package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 7, 11}
	b := a[2:]
	c := a[:2]
	for i := len(c); i > 0; i-- {
		b = append(b, c[i-1])
	}
	fmt.Println("部分反转：",b)

	reverse(a)

}

func reverse(s []int) {
	r := make([]int, 0)
	for i := len(s); i > 0; i-- {
		r = append(r, s[i-1])
	}
	fmt.Println("全部反转：",r)
}
