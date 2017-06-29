package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	str := "Welcome to Beijing, happy!"

	slice := strings.Fields(str)       //以空格键作为切割符,将str
	fmt.Println(reflect.TypeOf(slice)) //获取类型

	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println(slice)

}
