package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	if len(os.Args) != 3 {
		log.Fatal("bad command, example GET a.txt")
	}
	cmd, filename := os.Args[1], os.Args[2]
	conn.Write([]byte(cmd + " " + filename + "\n"))
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
	}
}
