package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	s := []int{2, 3, 1, 5, 9, 7}
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	fmt.Println(s)

	ss := []Student{}
	ss = append(ss, Student{
		Id:   1,
		Name: "aa",
	})

	ss = append(ss, Student{
		Id:   3,
		Name: "bb",
	})

	ss = append(ss, Student{
		Id:   2,
		Name: "cc",
	})

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Name < ss[j].Name
	})
	fmt.Println(ss)
}
