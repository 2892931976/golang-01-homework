package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Input Error! Please enter letters")
	} else {
		//var x byte
		str1 := os.Args[1]
		x := ""
		for _, s := range str1 {
			x = string(s) + x
		}
		fmt.Println(x)
	}
}
