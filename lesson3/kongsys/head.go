package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var lines = flag.Int("n", 10, "line")

func main() {
	// f, err := os.Create("a.txt")
	// f, err := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644);
	flag.Parse()
	unregargs := flag.Args()
	if len(unregargs) < 1 {
		fmt.Println("miss file.")
		return
	}
	f, err := os.Open(unregargs[0])
	if err != nil {
		log.Fatal(err)
	}

	// f.WriteString("hello\n")
	// f.Seek(1, os.SEEK_SET)
	// f.WriteString("$$")
	r := bufio.NewReader(f)
	for i := 0; i < *lines; i++ {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	f.Close()
}
