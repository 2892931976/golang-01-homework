package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	data := []byte(os.Args[1])
	md5sum := md5.Sum(data)
	fmt.Printf("%x %s\n", md5sum, os.Args[1])
	fmt.Printf("%#v\n", md5sum)

}
