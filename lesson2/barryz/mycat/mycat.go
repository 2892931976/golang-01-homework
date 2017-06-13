package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return
	}

	output := fmt.Sprintf(string(buf))

	fmt.Println(output)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("You must specify the a file name")
	}

	fileName := os.Args[1]
	printFile(fileName)
}
