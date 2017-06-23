package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var lines = flag.Int("n", 10, "display lines")

func main() {
	//读取文件就以Open打开就行
	f, err := os.Open(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	//逐行打印每一行

	r := bufio.NewReader(f)

	for i := 0; i < *lines; i++ {
		line, err := r.ReadString('\n')
		n, _ := strconv.Atoi(os.Args[2])
		if i == n || err == io.EOF {
			break
		}
		fmt.Print(line)
	}

	f.Close()

}
