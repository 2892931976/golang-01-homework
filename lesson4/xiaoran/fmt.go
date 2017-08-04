package main

import "fmt"

func main() {
	var s string
	var n int

	for {
		fmt.Print("> ")
		fmt.Scan(&s, &n)
		if s == "stop" {
			break
		}
		fmt.Println(s, n)
	}
}
