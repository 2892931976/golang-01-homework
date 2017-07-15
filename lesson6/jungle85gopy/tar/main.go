package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"./mytar"
)

func parseArgs(args []string) (dstType string, dstFile string, params []string) {
	if len(args) <= 2 {
		fmt.Printf("usage:\n\t%s xx.tar|yy.tar.gz files|paths ...\n", args[0])
		fmt.Printf("\ttar or tar.gz first, support multi files/paths\n")
		os.Exit(1)
	}
	dstFile = args[1]
	params = args[2:]
	if strings.HasSuffix(dstFile, mytar.DotTargz) {
		dstType = mytar.DotTargz
	} else if strings.HasSuffix(dstFile, mytar.DotTar) {
		dstType = mytar.DotTar
	} else {
		log.Fatalf("first params: %s with error type\n", dstFile)
	}
	return
}

func main() {
	dstType, dstFile, params := parseArgs(os.Args)
	fmt.Printf("main:\n\ttype: %s, dest: %s. params: %v\n\n", dstType, dstFile, params)

	err := mytar.MyTar(dstType, dstFile, params)
	if err != nil {
		log.Println(err)
	}
}
