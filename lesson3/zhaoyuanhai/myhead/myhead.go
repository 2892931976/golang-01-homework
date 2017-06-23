package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func readFile(f string, n int) {
	fh, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(fh)

	for i := 0; i < n; i++ {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
	fh.Close()
}

func main() {
	if len(os.Args) == 2 {
		n := 10
		readFile(os.Args[1], n)
	} else if len(os.Args) == 4 && os.Args[1] == "-n" {
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("输入的行数不正确 head -n 10 filename")
			return
		}
		readFile(os.Args[3], n)
	} else {
		fmt.Println("输入参数不正确")
	}

}
