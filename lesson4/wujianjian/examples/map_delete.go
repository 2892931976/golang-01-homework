package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
	}
	fmt.Println(m["a"])
	delete(m, "a")
	fmt.Println(m["a"])

	//空map是nil，是不能使用的，需要make初始化才可以使用
	var m1 map[string]int
	if m1 == nil {
		fmt.Println("m1 map is null")
	}
	fmt.Println(m1)
	m1 = make(map[string]int)
	fmt.Println(m1)
}
