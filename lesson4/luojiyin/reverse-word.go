package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world 欢 迎 入 坑 golang"
	fmt.Printf("Fileds are :%q\n", strings.Fields(s))
	reverseWord(s)
}

func reverseWord(s string) (reverse []string) {
	temp := strings.Fields(s)

	for i := len(temp) - 1; i >= 0; i-- {
		//fmt.Printf("%v\n", temp[i])
		reverse = append(reverse, temp[i])
	}
	fmt.Println(reverse)
	return reverse

}
