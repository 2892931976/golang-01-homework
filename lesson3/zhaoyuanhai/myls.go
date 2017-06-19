package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// f, err := os.Open(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// infos, _ := f.Readdir(-1)
	// for _, info := range infos {
	// 	fmt.Printf("%v %d %s\n", info.IsDir(), info.Size(), info.Name())
	// }
	// f.Close()
	fname := os.Args[1]
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	infos, _ := f.Readdir(-1)
	for _, info := range infos {
		fmt.Printf("%v %d %s\n", info.IsDir(), info.Size(), info.Name())
	}
	f.Close()

}
