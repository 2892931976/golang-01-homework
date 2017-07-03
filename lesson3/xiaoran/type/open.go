package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0644)
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)
	for i := 0; i < 10; i++ { //循环读取10行
		line, err := r.ReadString('\n') //遇到\n 读一行，也可以换成空格或制表符\t
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	f.Close()
}
