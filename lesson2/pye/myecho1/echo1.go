package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep1 = flag.String("s", " ", "separator")
var sep2 = flag.String("q", "|", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep1))
	fmt.Println(strings.Join(flag.Args(), *sep2))
}
