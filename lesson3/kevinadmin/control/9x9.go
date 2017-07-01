package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, error := os.Create("n9n.txt")
	if error != nil {
		fmt.Println("error is ", error)
		return
	}
	for i := 1; i < 10; i++ {
		for k := 1; k <= i; k++ {
			fmt.Fprintf(file, "%dx%d=%-2d ", k, i, i*k)
		}
		fmt.Fprintf(file, "\n")
	}
	file.Close()

	f, err := os.Open("n9n.txt")
	if err != nil {
		fmt.Println("it is a error ")
		return
	}

	fr := bufio.NewReader(f)
	for {
		s, readError := fr.ReadString('\n')
		fmt.Println(s)
		if readError == io.EOF {
			return
		}
	}
}
