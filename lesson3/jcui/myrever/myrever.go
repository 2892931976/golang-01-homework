package main

import (
	"fmt"
	"os"
)

func reverser(s string) string {
	result := []byte(s)
	for x := 0; x < len(result)-1; x++ {
		for y := x + 1; y < len(result); y++ {
			//fmt.Println(result[x])
			//fmt.Println(result[y])
			//fmt.Println("")
			result[x], result[y] = result[y], result[x]
		}
	}
	return string(result)
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("args error")
		os.Exit(0)
	}
	mystr := os.Args[1]
	fmt.Println(reverser(mystr))
}
