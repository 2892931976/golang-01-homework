package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	conn.Write([]byte("STORE /etc/passwd\n"))
	io.Copy(conn, os.Stdin)
}
