/*
http://www.jianshu.com/p/172810a70fad
*/
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

func readAddr(r *bufio.Reader) (string, error) {
	version, _ := r.ReadByte()
	log.Printf("version:%d", version)
	if version != 5 {
		return "", errors.New("bad version")
	}

	cmd, _ := r.ReadByte()
	log.Printf("cmd:%d", cmd)

	r.ReadByte()

	addratyp, _ := r.ReadByte()
	log.Printf("addratyp:%d", addratyp)
	if addratyp != 3 {
		return "", errors.New("bad addratyp")
	}

	// addrlen
	addrlen, _ := r.ReadByte()
	log.Printf("addrlen:%d", addrlen)

	// addr
	addr := make([]byte, addrlen)
	io.ReadFull(r, addr)
	log.Printf("addr:%s", addr)

	var port int16
	binary.Read(r, binary.BigEndian, &port)
	return fmt.Sprintf("%s:%d", addr, port), nil
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

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	handshake(r, conn)
	addr, _ := readAddr(r)
	log.Printf("addr:%s", addr)

	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)

	// 开始代理
	//server, err := net.Dial("tcp", "www.baidu.com:80")
	server, err := net.Dial("tcp", addr)
	if err != nil {
		log.Print(err)
		return
	}
	defer server.Close()
	go io.Copy(server, conn)
	io.Copy(conn, server)

}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", ":8022")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		log.Print("wait connect.")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("new %s", conn)
		go handleConn(conn)
	}

}
