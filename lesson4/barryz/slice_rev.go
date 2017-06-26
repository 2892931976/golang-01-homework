package main

import (
	"fmt"
	"os"
)

func main() {
	a := []int{2, 3, 5, 6, 7}
	target := 1
	s, err := sliceJoin(a, target)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(s)
}

func sliceJoin(s []int, target int) (sr []int, err error) {
	if target > len(s) {
		err = fmt.Errorf("target greater than the params's length.")
		return
	}

	s1, s2 := s[0:target], s[target:]
	sr = append(s2, s1...)
	return
}
