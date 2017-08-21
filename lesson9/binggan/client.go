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
	// connection
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr().String())
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("write size", n)

	// buf := make([]byte, 10)
	// n, err = conn.Read(buf)
	// if err != nil && err != io.EOF {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(buf[:n]))
	io.Copy(os.Stdout, conn)
	conn.Close()
}
