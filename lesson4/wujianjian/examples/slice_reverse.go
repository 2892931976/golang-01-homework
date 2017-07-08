package main

import "fmt"

//反转切片
func reverse(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}

func main() {
	s := []int{2, 3, 5, 7, 11}
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
	reverse(s[1:4])
	fmt.Println(s)

}
