package main

import (
	"io"
	"log"
	"os"
)

func v1() {
	var f *os.File
	var err error
	if len(os.Args) > 1 {
		f, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f = os.Stdin
	}

	buf := make([]byte, 1024)
	for {
		//n, err := os.Stdin.Read(buf)
		n, err := f.Read(buf)
		if err != nil {
			return
		}
		os.Stdout.Write(buf[:n])
	}
}

func v2() {
	var f *os.File
	var err error
	if len(os.Args) > 1 {
		f, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f = os.Stdin
	}
	//把右边的数据拷贝到左边
	io.Copy(os.Stdout, f)

}

func main() {
	//v1()
	v2()
}
