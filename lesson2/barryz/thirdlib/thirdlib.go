package main

import (
	"fmt"
	"os"

	"github.com/51reboot/golang-01-homework/lesson2/barryz/thirdlib/xmath"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(xmath.Add(2, 3))

	fmt.Println(s)
}
