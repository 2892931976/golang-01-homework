package main

import (
	"os"
	"fmt"
	"kevinadmin.org/golib"
	golib2 "github.com/kevinadmin/golib"
)

func main (){
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(X)
	fmt.Println(golib.Add(1, 3))
	fmt.Println(golib2.Adds(1, 5))
}