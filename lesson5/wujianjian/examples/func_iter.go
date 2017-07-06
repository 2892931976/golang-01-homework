package main

import (
	"errors"
	"fmt"
)

// 函数的返回类型是一个闭包
// 闭包：就是基本的函数，去掉了函数名，剩下的部分，有匿名函数的地方，必然有闭包
func iter(s []int) func() (int, error) {
	var i = 0
	return func() (int, error) {
		if i >= len(s) {
			return 0, errors.New("end")
		}
		n := s[i]
		i += 1
		return n, nil
	}
}

func main() {
	f := iter([]int{1, 2, 3})
	for {
		n, err := f()
		if err != nil {
			break
		}
		fmt.Println(n)
	}
}
