package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "abc123"
	if n, err := strconv.Atoi(s); err == nil {
		fmt.Println(n)
	}
	switch s {
	case "1":
		fmt.Println("s=", "1")
	case "2":
		fmt.Println("s=", "2")

	default:
		fmt.Println("s=", "default", s)
	}
}
