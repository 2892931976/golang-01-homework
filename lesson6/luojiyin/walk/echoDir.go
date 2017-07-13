package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func printDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	if info.IsDir() {
		fmt.Println(path)
	}
	return nil
}

func main() {
	log.SetFlags(log.Lshortfile)
	dir := os.Args[1]
	err := filepath.Walk(dir, printDir)
	if err != nil {
		log.Fatal(err)
	}
}
