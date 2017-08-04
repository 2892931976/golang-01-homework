package main

import "fmt"

func main() {
	//空的struct是不占用空间的
	//set := make(map[string]bool)
	set := make(map[string]struct{})
	set["a"] = struct{}{}
	set["a"] = struct{}{}
	if _, ok := set["b"]; ok {
		fmt.Println("ok")
	} else {
		fmt.Println("!ok")
	}
}
