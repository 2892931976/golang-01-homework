package main

import (
	"fmt"
	"strings"
)

func main() {
	// 翻转切片元素,结果为：[44 33 22 11 7 5 3 2]
	s := []int{2, 3, 5, 7, 11, 22, 33, 44}
	fmt.Println(reverse(s))

	//翻转切片
	s1 := []int{2, 3, 5, 7, 11, 22, 33, 44}
	//翻转前段切片元素，前段切片元素顺序为：[5,3,2]；s1的切片顺序为：[5 3 2 7 11 22 33 44]
	reverse(s1[:3])
	//fmt.Println(s1)

	//翻转后段切片元素，后段切片元素顺序为：[44,33,22,11,7]；s1的切片顺序为：[5 3 2 44 33 22 11 7]
	reverse(s1[3:])
	//fmt.Println(s1)

	//再对s1现有顺序进行翻转，得出想要的结果：[7 11 22 33 44 2 3 5]
	fmt.Println(reverse(s1))

	//翻转单词
	ss := "hello world ggg"
	fmt.Println(reverseString(ss))
}

func reverse(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	//fmt.Println(s)
	return s
}

func reverseString(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	words1 := strings.Join(words, " ")
	return words1

}
