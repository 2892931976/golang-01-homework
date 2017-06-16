package main

import (
	"flag"
	"fmt"
	"strings"
)

//第一个参数指 s 代表 -s
//第二个参数代表默认值，这里默认为空
//第三个参数显示（一个说明）
var sep = flag.String("s", " ", "separator")
var newline = flag.Bool("n", false, "newline")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *newline {
		fmt.Println()
	}
}
