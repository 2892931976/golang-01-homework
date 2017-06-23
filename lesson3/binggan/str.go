package main

import "fmt"

func main() {
	s1 := "hello" + "world"

	s2 := "helloworld"

	if s1 == s2 {
		fmt.Println("equal")
	}

	fmt.Println(0, len(s1)-1)

	var c1 byte
	c1 = s1[0]
	fmt.Println(s1, s2, c1)
	fmt.Printf("数字:%d 字符:%c\n", c1, c1)

	s3 := s1[:]
	fmt.Println(s3)

	var b byte
	for b = 0; b < 177; b++ {
		fmt.Printf("%d %c\n", b, b)
	}

	array := []byte(s1)
	fmt.Println(array)
	array[0] = 72 // 'H'
	s1 = string(array)
	fmt.Println(s1)

	fmt.Println('a' + ('H' - 'h'))

	fmt.Println(0xa)
	fmt.Println(toupper("hello"))
}

func toupper(s string) string {
	// len(s)
	// fmt.Println('a' + ('H' - 'h'))
	// array := []byte(s1)
	var result string
	return result
}
