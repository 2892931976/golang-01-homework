//
// fibonacci-closure.go
// Copyright (C) 2017 root <root@localhost.localdomain>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

func fibonacci() func(int64) int64 {
	var num1, num2 int64
	num1, num2 = 1, 1
	return func(x int64) int64 {
		if x < 2 {
			return num2
		}
		temp := num1 + num2
		num1, num2 = num2, temp
		return temp
	}
}

func main() {
	pos := fibonacci()

	for i := 0; i < 94; i++ {
		fmt.Printf("%d\n", pos(int64(i)))
	}
}
