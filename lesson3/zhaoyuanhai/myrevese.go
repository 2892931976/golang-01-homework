package main

import (
	"fmt"
	//	"os"
)

func reverse1(s string) {
	var s2 string
	s1 := []rune(s)
	for i := len(s1) - 1; i >= 0; i-- {
		s2 = s2 + string(s1[i])
	}
	fmt.Println(s2)
}

func reverse2(s string) {
	s3 := []rune(s)
	for i := 0; i <= len(s3)/2; i++ {
		s3[i], s3[len(s3)-1-i] = s3[len(s3)-1-i], s3[i]
	}
	fmt.Println(string(s3))
}

func main() {
	// if len(os.Args) > 1 {
	// 	//		var s string = os.Args[1]
	// 	reverse("abc中国aaa")
	// } else {
	// 	fmt.Println("输入参数不正确(myreverse stringName)")
	// }
	reverse1("abc中国aaa")
	reverse2("abc中国aaab")
}
