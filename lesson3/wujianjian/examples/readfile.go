package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		log.Fatal(err)
	}

	//申请1024字节的内存空间
	//buf := make([]byte, 1024)
	//f.Read(buf)

	//读取一行，bufio是go的函数库
	/*
		r := bufio.NewReader(f)
		line, _ := r.ReadString('\n')
		fmt.Print(string(line))
		line, _ = r.ReadString('\n')
		fmt.Print(string(line))
	*/
	//fmt.Println(string(buf))

	//逐行打印每一行
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}

	//随机读取
	f.Seek(3, os.SEEK_SET)
	buf := make([]byte, 2)
	f.Read(buf)
	fmt.Println(string(buf))

	f.Close()
}
