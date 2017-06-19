package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", " ", "separator")
var newline = flag.Bool("n", false, "是否换行,类似echo里的 -n")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *newline {
		fmt.Println()
	}
}
