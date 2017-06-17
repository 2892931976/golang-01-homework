package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("table.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprintf(f, "%dx%d=%-2d ", i, j, i*j)
		}
		fmt.Fprintf(f, "\n")
	}
	f.Close()
}
