package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
	// tmps()
}

func tmps() {
	var (
		x int
		y float32
		z string
		p *int
		b bool
	)
	i := 0
	s := "hello"
	j, k := 0, 1
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", s)
	fmt.Println(k, j)
}
