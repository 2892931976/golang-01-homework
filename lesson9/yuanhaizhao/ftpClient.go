package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("bad command,example ftp 10.13.0.0 22")
	}
	for {
		conn, err := net.Dial("tcp", os.Args[1]+":"+os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		var cmd, filename string
		f := bufio.NewReader(os.Stdin)
		fmt.Printf("%v", "> ")
		line, _ := f.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Sscan(line, &cmd, &filename)
		// if len(os.Args) != 3 {
		// 	log.Fatal("bad command, example GET a.txt")
		// }
		conn.Write([]byte(cmd + " " + filename + "\n"))
		fmt.Println(cmd, filename)
		switch cmd {
		case "GET":
			f, err := os.Create(filename)
			if err != nil {
				log.Print(err)
				return
			}
			io.Copy(f, conn)
		case "STORE":
			f, err := os.Open(filename)
			if err != nil {
				log.Print(err)
				return
			}
			io.Copy(conn, f)
		case "LS":
			io.Copy(os.Stdout, conn)
		case "quit":
			os.Exit(0)
		}
	}
}
