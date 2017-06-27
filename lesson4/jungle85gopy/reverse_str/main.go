package main

import (
	"fmt"
	"strings"
)

func revStr(s string) string {
	var rst string
	var temp = strings.Fields(s)
	fmt.Println("temp rst :", temp)
	for i := len(temp) - 1; i >= 0; i-- {
		rst = rst + temp[i] + " "
	}
	return rst
}

func main() {
	fmt.Println("reverse by words in strings:")
	var strs = "hello 		world   for my	 golang"
	fmt.Println("original :", strs)
	fmt.Println("reversed :", revStr(strs))
}
