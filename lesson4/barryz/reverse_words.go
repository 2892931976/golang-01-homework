package main

import (
	"fmt"
	"strings"
)

func main() {
	testStr := "The deepest shade of blue is waking up by your side"
	testStrSlice := strings.Split(testStr, " ")

	for i := 0; i < (len(testStrSlice) >> 1); i++ {
		mI := len(testStrSlice) - i - 1
		testStrSlice[i], testStrSlice[mI] = testStrSlice[mI], testStrSlice[i]
	}

	testStrRev := strings.Join(testStrSlice, " ")
	fmt.Println(testStrRev)
	// output: side your by up waking is blue of shade deepest The
}
