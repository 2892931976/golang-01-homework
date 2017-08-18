package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr := "www.baidu.com:80"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr().String())
	// n, err := conn.Write([]byte("GET / HTTP/1.1\r\nConnection: Close\r\n\r\n"))
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("write size", n)
	io.Copy(os.Stdout, conn)
	conn.Close()
}
