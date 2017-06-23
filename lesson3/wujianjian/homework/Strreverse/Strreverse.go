package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	str = os.Args[1]
	//fmt.Println(str)
	array := []byte(str)
	restr := []byte(str)
	//fmt.Println(array)
	for i := len(array) - 1; i >= 0; i-- {
		restr[len(array)-1-i] = array[i]
	}
	fmt.Println(string(restr))
}
