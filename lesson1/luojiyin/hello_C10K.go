package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handle(conn net.Conn) {
	fmt.Fprint(conn, "%s", time.Now().String())
	conn.Close()
}

func main() {
	lister, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lister.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}
