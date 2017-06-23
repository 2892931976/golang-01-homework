package main

import "fmt"

func toReverser(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	s := "hello  world"
	fmt.Println(s)
	fmt.Println(toReverser(s))

}
