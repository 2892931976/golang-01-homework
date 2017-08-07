package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8021")
	if err != nil {
		log.Fatal(err)
	}

	conn.Write([]byte)
}

func handleConn(conn net.Conn) {
	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')

	var content []byte
	conn.Write(content)

	//-----------------------
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	header := buf[:n]

	var content []byte
	conn.Write(content)
}
