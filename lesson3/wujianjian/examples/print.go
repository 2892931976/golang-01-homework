package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//打印格式写入文件中去
	//第一步：现在屏幕打印出来
	//第二步：写到文件中去
	f, err := os.Create("fmt.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(f, "hello")
	fmt.Fprintln(f, "hello")
	s := "hello"
	n := 4
	fmt.Fprintf(f, "my string is: %s n=%d\n", s, n)
	f.Close()
}
