package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("99.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprintf(f, "%d * %d = %2d  ", j, i, j*i)

		}
		fmt.Fprintln(f)
	}
}
