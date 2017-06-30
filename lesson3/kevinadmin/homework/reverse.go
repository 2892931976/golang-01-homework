package main

import (
	"fmt"
)

func main() {

	fmt.Println(reverse("hello world"))
	fmt.Println("Other --> ", reverseTwo("is is to here"))

}

func reverse(s string) string {
	fmt.Println("Start to Reverse ---> ", s)
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		if ss[i] > 'a' && ss[i] < 'z' {
			ss[i] -= 32
		}
	}
	return string(ss)
}

func reverseTwo(s string) string {
	rs := []rune(s)
	lens := len(rs)
	var tt []rune

	tt = make([]rune, 0)
	for i := 0; i < lens; i++ {
		tt = append(tt, rs[lens-i-1])
	}
	return string(tt)
}
