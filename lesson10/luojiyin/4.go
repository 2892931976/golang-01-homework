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

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	handershake(r, conn)
	readAddr(r)
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)
}

func readAddr(r *bufio.Reader) (string, error) {
	version, _ := r.ReadByte()
	log.Printf("version:%d", version)
	if version != 5 {
		return
	}

	cmd, _ := r.ReadByte()
	if cmd != '1' {
		return
	}
	//跳过RSV
	r.ReadByte()
	atyp, _ := r.ReadByte()
	log.Printf("addr type:%d", addrtype)
	if atyp != 3 {
		return "", errors.New("bad addr type ")
	}

	addrlen, _ := r.ReadByte()

	addr := make([]byte, addrlen)
	io.ReadFull(r, addr)

	var port int16
	binary.Read(r, binary.BigEndian, &port)
	return fmt.Sprintf("%s:%d", addr, port), nil

	//bnd_port ,_ := r.ReadByte()
}

func handshake(r *bufio.Reader, conn net.Conn) error {
	version, _ := r.ReadByte()
	log.Printf("version:%d", version)
	if version != 5 {
		return errors.New("bad version")
	}
	nmethods, _ := r.ReadByte()
	log.Printf("nmethods:%d", nmethods)

	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	log.Printf("%v", buf)

	resp := []byte{5, 0}

	conn.Write(resp)
	return nil
}
func main() {
	flag.Parse()
	l, err := net.Listen("tcp", "8021")
	if err != nil {
		log.Print(err)
	}
	for {
		conn, _ := l.Accept()
	}

}
