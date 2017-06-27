package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", "", "seperator")
var Newline = flag.Bool("n", false, "new line")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *Newline {
		fmt.Println()
	}
}
