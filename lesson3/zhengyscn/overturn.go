package main

import "fmt"

func reverseString(str string) string {
	n := len(string(str[:])) - 1
	var runes string
	for {
		if n < 0 {
			break
		}
		runes = runes + string(str[n])
		n--
	}
	return runes
}

func main() {
	x := "hello world"
	fmt.Println(reverseString(x))

}
