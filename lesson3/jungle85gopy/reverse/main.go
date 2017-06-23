package main

import "fmt"

// Reverse func reverse a string
func Reverse(origin string) string {
	fmt.Println(origin)
	sLen := len(origin)
	if sLen == 0 {
		return origin
	}
	buf := []byte(origin)
	for sta, mid, end := 0, sLen/2, sLen-1; sta < mid; sta++ {
		buf[sta], buf[end] = buf[end], buf[sta]
		// sta++
		end--
	}
	return string(buf)
}

func main() {
	testStr := "hello golang"
	fmt.Println(Reverse(testStr))
}
