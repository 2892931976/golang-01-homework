package main

import (
	"fmt"
	"net"

	"flag"
	"io"
	"log"
	"os"
)

var (
	filename = flag.String("filename", "", "upload local file to server.")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", "192.168.1.103:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// dd if=/dev/zero bs=1M count=1024 of=file
	cmd := fmt.Sprintf("STORE %s\n", *filename)
	conn.Write([]byte(cmd))
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(conn, f)
}
