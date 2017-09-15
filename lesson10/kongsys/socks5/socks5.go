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
	"time"
)

var timeout = flag.Int("t", 5, "connect timeout, default 5s.")

func mustReadByte(r *bufio.Reader) byte {
	b, err := r.ReadByte()
	if err != nil {
		panic(err)
	}

	return b
}
func readAddr(r *bufio.Reader) (addr string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error) // interface{}
		}
	}()
	// version, _ := r.ReadByte()
	version := mustReadByte(r)
	// log.Printf("readAddr version: %d", version)
	if version != 5 {
		return "", errors.New("bad Version")
	}
	// cmd, _ := r.ReadByte()
	cmd := mustReadByte(r)
	switch cmd {
	case 1:
		log.Printf("readAddr cmd: %d, CONNECT", cmd)
	case 2:
		log.Printf("readAddr cmd: %d, BIND", cmd)
	case 3:
		log.Printf("readAddr cmd: %d, UDP", cmd)
	default:
		log.Printf("readAddr cmd: %d, unknow", cmd)
		return "", errors.New("err cmd")
	}

	mustReadByte(r)
	atyp := mustReadByte(r)
	var addrlen int
	switch atyp {
	case 1:
		// log.Printf("readAddr atyp: %d, ipV4", atyp)
		addrlen = 4
	case 3:
		// log.Printf("readAddr atyp: %d, domain", atyp)
		addrl, _ := r.ReadByte()
		// fmt.Println(addrl)
		addrlen = int(addrl)
	case 4:
		// log.Printf("readAddr atyp: %d, ipV6", atyp)
		addrlen = 16
	}
	host := make([]byte, addrlen)
	var hh string
	io.ReadFull(r, host)
	// fmt.Printf("host:%s\n", host)
	switch atyp {
	case 1:
		hh = net.IPv4(host[0], host[1], host[2], host[3]).String()
	case 3:
		hh = string(host)
		// log.Printf("hh:%s", hh)
	}
	// log.Printf("readAddr domain: %s", hh)

	var port int16
	binary.Read(r, binary.BigEndian, &port)

	return fmt.Sprintf("%s:%d", hh, port), nil
}

func handshake(r *bufio.Reader, conn net.Conn) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error) // interface{}
		}
	}()
	version := mustReadByte(r)
	// log.Printf("version:%d", version)
	if version != 5 {
		return errors.New("bad Version")
	}

	nmethods := mustReadByte(r)
	// log.Printf("nmethods:%d", nmethods)

	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	// log.Printf("%v", buf)

	resp := []byte{5, 0}
	// log.Printf("resp %T\n", resp)
	conn.Write(resp)
	return nil
}

func handleConn(conn net.Conn) {
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
		}
	}()
	defer conn.Close()

	r := bufio.NewReader(conn)
	err := handshake(r, conn)
	if err != nil {
		log.Println(err)
		return
	}
	addr, err := readAddr(r)
	if err != nil {
		log.Println(err)
		return
	}
	// log.Printf("addr: %s", string(addr))
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)
	server, err := net.DialTimeout("tcp", addr, time.Duration(*timeout)*time.Second)
	if err != nil {
		log.Println(err)
		return
	}

	defer server.Close()

	s := bufio.NewReader(server)
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		n, err := io.Copy(server, r)
		log.Println("func io copy end", n, err)
		conn.SetDeadline(time.Now())
		server.SetDeadline(time.Now())
	}()

	n, err := io.Copy(conn, s)
	log.Println("func io copy end", n, err)
	conn.SetDeadline(time.Now())
	server.SetDeadline(time.Now())
}

func main() {
	//flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":8022")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
