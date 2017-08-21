package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

var (
	target = flag.String("target", "www.baidu.com", "target host")
)

func handleConn(conn net.Conn) {
	var remote net.Conn
	remoteaddr := fmt.Sprintf("%s:80", *target)
	defer conn.Close()
	remote, err := net.Dial("tcp", remoteaddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer remote.Close()
	go io.Copy(remote, conn)
	io.Copy(conn, remote)
}

func main() {
	listen, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
