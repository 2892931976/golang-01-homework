package main

import (
	"fmt"
	"os"
)

func myreverse(s string) string {
	var zifu string
	y := []byte(s)

	// fmt.Println(y)
	for i := len(y) - 1; i >= 0; i-- {

		zifu = zifu + string(y[i])

	}
	return string(zifu)

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("you should input args 1 ")
	} else {
		zh := os.Args[1]
		fmt.Println(myreverse(zh))

	}

}
