package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	s := os.Args[1]
	fmt.Println(s)
	f, err1 := os.Open(s)
	if err1 != nil {
		fmt.Println("error")
		log.Fatal(err1)
		os.Exit(0)
	}
	infos, _ := f.Readdir(-1)
	for _, info := range infos {
		if info.IsDir() {
			fmt.Printf("%v %d %s\n", info.IsDir(), info.Size(), info.Name())
		}
	}
	f.Close()

}
