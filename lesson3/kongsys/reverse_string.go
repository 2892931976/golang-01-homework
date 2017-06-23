package main

import (
	"fmt"
	"github.com/51reboot/golang-01-homework/lesson3/kongsys/reverse_lib"
)

func main() {
	var input string = "你好世界，just demo"
	// var input string = "just demo"
	fmt.Println("origin :", input)
	re_input := reverse.Reverse(input)
	fmt.Println("reverse :", re_input)
}
