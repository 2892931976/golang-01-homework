package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Id   int
}

func main() {
	s := []int{2, 3, 1, 5, 9, 7}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	ss := []Student{}
	ss = append(ss, Student{
		Name: "aa",
		Id:   2,
	})
	ss = append(ss, Student{
		Name: "cc",
		Id:   3,
	})
	ss = append(ss, Student{
		Name: "bb",
		Id:   1,
	})

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Name < ss[j].Name
	})

	fmt.Println(ss)
}
