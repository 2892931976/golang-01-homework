package main

import (
	"fmt"
	"strings"
)

func reverse(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}

func main() {
	/*
		题1：
		初始数据：2 3 5 7 11
		目标数据：5 7 11 2 3
		第一步：3 2 ,第二步：11 7 5，第三步：连接后再反转
	*/
	s := []int{2, 3, 5, 7, 11}

	s1 := s[0 : len(s)/2]
	s2 := s[len(s)/2:]
	s1 = reverse(s1)
	s2 = reverse(s2)
	s1 = append(s1, s2...)
	s = reverse(s1)
	fmt.Println(s)

	/*
		题2：
		原始数据：hello world
		目标数据：world hello
	*/

	str := "hello world"
	var res []string
	res = strings.Split(str, " ")
	var t = make([]string, len(res))
	for i := 0; i < len(res); i++ {
		t[i] = res[len(res)-i-1]
	}
	fmt.Println(t)
}
