package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	//切片是左闭右开的
	var s []int = primes[1:4]
	fmt.Println(s)

}
