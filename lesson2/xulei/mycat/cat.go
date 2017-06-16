package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func printFile(name string) {
	fmt.Println(name)
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}

func main() {
	flag.Parse()
	for i := 0; i < len(flag.Args()); i++ {
		// fmt.Println(len(flag.Args()))
		printFile(flag.Arg(i))
	}
}
