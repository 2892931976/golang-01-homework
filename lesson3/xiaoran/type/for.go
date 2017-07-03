package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 2; i < 3; i++ {
		fmt.Println(i)
	}

	i := 5
	for i < 7 {
		fmt.Println(i)
		i = i + 1
	}

	i = 8
	for {
		i = i + 1
		fmt.Println(i)
		if i > 10 {
			break
		}
	}

	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	s := "hello中文"
	for i, arg := range s {
		fmt.Printf("%d %c\n", i, arg)
	}
}
