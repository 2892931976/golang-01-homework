package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//打印9X9乘法表，并写入文件
	f, err := os.Create("fmt.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Fprintf(f, "%d * %d = %d", j, i, j*i)
			fmt.Fprint(f, "  ")
		}
		fmt.Fprintln(f, "")
	}
	f.Close()

	//读取9X9乘法表文件内容，并打印出来
	fh, err := os.Open("fmt.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(fh)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
	fh.Close()

	// fd, err := os.OpenFile("fmt.txt", os.O_RDONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fd.Seek(0, os.SEEK_END)
	// r := bufio.NewReader(fh)

	// for {
	// 	line, err := r.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(line)
	// }
	// fh.Close()
	// fd, err := os.Open("fmt.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fd.Seek(-1, os.SEEK_END)
}
