package main

import "fmt"

func myreverse(oldstr string) {
	newstr := []rune(oldstr)
	for i := len(newstr) - 1; i >= 0; i-- {
		fmt.Printf(string(newstr[i]))
	}

}

func main() {
	myreverse("abcdeeqwewqqwewqfg")
}
