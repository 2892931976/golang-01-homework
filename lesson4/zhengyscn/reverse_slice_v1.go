package main

import "fmt"

func ReversePos(n int, s []int) {
	s = append(s[n:], s[:n]...)
	fmt.Println(s)
}

func main() {
	s := []int{2, 3, 5, 7, 11}
	ReversePos(4, s)

}
