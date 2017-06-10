package main

import (
	"fmt"
	"os"
  "github.com/51reboot/golang-01-homework/lesson2/kongsys/myecho/mathlib"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
  fmt.Println(mathlib.Add(3, 5))
}
