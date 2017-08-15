package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

var (
	listen = flag.String("l", ":9005", "listen addr")
)

func readAddr(r *bufio.Reader) (string, error) {
	version, _ := r.ReadByte()
	log.Printf("version: %d", version)
	if version != 5 {
		return "", errors.New("bad version")
	}
	cmd, _ := r.ReadByte()
	log.Printf("cmd: %d", cmd)
	if cmd != 1 {
		return "", errors.New("bad cmd")
	}
	r.ReadByte()
	addrtype, _ := r.ReadByte()
	log.Printf("addrtype: %d", addrtype)
	if addrtype != 3 {
		return "", errors.New("bad addr")
	}
	addrlen, _ := r.ReadByte()
	log.Printf("addrlen: %d", addrlen)
	addr := make([]byte, addrlen)
	io.ReadFull(r, addr)
	log.Printf("%s", addr)
	var port int16
	binary.Read(r, binary.BigEndian, &port)
	return fmt.Sprintf("%s:%d", addr, port), nil
}

func handshake(r *bufio.Reader, conn net.Conn) error {
	version, _ := r.ReadByte()
	log.Printf("version: %d", version)
	if version != 5 {
		return errors.New("bad version")
	}
	nmethods, _ := r.ReadByte()
	log.Printf("nmethods: %d", nmethods)
	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	log.Printf("%s", buf)
	resp := []byte{5, 0}
	conn.Write(resp)
	return nil
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	handshake(r, conn)
	addr, _ := readAddr(r)
	log.Printf("addr: %v", addr)
	resp := []byte{0x05, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)
	remote, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	defer remote.Close()
	go io.Copy(remote, r)
	io.Copy(conn, remote)
}

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConn(conn)
	}
}
