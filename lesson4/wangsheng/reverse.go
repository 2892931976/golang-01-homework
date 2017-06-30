package main

import (
	"fmt"
	//"strings"
)

func rever_slice(a []int, x int) []int {

	//传入的slice

	result := a[x:]
	//fmt.Println(result)
	for j := 0; j < x; j++ {
		result = append(result, a[j])
	}
	return result
}

func main() {
	a := []int{2, 3, 5, 7, 11}
	//[2, 3, 5, 7, 11]
	fmt.Println(a)

	i := 3

	///fmt.Println(i)

	b := rever_slice(a, i)
	fmt.Println(b)

}
