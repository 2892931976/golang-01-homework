package main

import "fmt"

func main() {
	str := "abCdefg中文"
	fmt.Printf("str (rune array):	%v\n", []rune(str))

	fmt.Printf("Raw string:\n%s\n", `a\t
	b`)
}
