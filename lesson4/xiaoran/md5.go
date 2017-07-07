package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	str := os.Args[1]
	buf, err := ioutil.ReadFile(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := []byte(string(buf))
	md5sum := md5.Sum(data)
	fmt.Println(md5sum)
	fmt.Printf("%x\n", md5sum)
}
