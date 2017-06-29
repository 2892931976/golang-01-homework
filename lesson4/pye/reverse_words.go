package main

import (
	"fmt"
	"strings"
)

func main() {
	twords := "hello this world let us rock"
	u := strings.Split(twords, " ")
	a := make([]string, 0, len(u))
	for i := len(u) - 1; i >= 0; i-- {
		a = append(a, u[i])
	}

	fmt.Println(twords)
	fmt.Println(strings.Join(a, " "))
}
