package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func myslice(s []int, n int) error {
	if n < 0 {
		return errors.New("截断长度小于0")
	}

	if n == 0 || n >= len(s) {
		return nil
	}

	var t []int
	t = append(t, s[n:]...)
	t = append(t, s[:n]...)
	copy(s, t)
	return nil
}

func mystring(str string) string {
	s := strings.Fields(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return strings.Join(s, " ")
}

func main() {
	a := []int{2, 3, 5, 7, 11}
	fmt.Println(a)
	err := myslice(a, 2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(a)
	fmt.Println()

	str := "hello world"
	fmt.Println(str)
	fmt.Println(mystring(str))
}
