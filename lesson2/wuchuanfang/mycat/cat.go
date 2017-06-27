package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}

func main() {
	//var s string
	for i := 1; i < len(os.Args); i++ {
		printFile(os.Args[i])
		//s += sep + os.Args[i]
		//sep = " "
	}

}
