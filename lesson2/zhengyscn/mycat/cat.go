package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var numberFlag = flag.Bool("n", false, "show number")

func printFile(filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	if *numberFlag {
		data_list := strings.Split(string(buf), "\n")
		for i, v := range data_list {
			if len(v) != 0 {
				fmt.Printf("    %v  %v", i+1, v)
				fmt.Println()
			}
		}
	} else {
		fmt.Println(string(buf))
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%v filename", os.Args[0])
		fmt.Println()
		return
	}

	flag.Parse()
	for i := 0; i < len(flag.Args()); i++ {
		printFile(flag.Arg(i))
	}

}
