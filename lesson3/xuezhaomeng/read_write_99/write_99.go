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

	for i := 1; i < 10; i++ {
		for ii := 1; ii <= i; ii++ {
			fmt.Fprintf(f, "%d * %d = %d ", ii, i, i*ii)
		}
		fmt.Fprintln(f, "")

	}
	f.Close()

}
