package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSice(s)

	s = s[:0]
	printSice(s)

	s = s[:4]
	printSice(s)

	s = s[2:]
	printSice(s)

}

func printSice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
