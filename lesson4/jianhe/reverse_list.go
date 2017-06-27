package main

import (
	"fmt"
	"os"
)

func reverse(lst []int, n int) chan int {

	if n >= len(lst) {
		// return nil   这句话会造成deadlock！  为什么？？？？？？
		os.Exit(1)
	}
	ret := make(chan int)
	go func() {
		for i, _ := range lst {
			ret <- lst[n+i]
			if n+i == len(lst)-1 {
				break
			}
		}
		for i := 0; i < n; i++ {
			ret <- lst[i]
		}
		close(ret)
	}()
	return ret
}

func main() {
	elms := []int{1, 3, 4, 2, 8, 7, 6}
	for e := range reverse(elms, 4) {
		fmt.Println(e)
	}
}
