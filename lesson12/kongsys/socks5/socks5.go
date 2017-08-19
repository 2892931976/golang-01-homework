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

	"github.com/51reboot/golang-01-homework/lesson11/kongsys/mycrypto"
)

var (
	key    = flag.String("key", "123456", "cipher key")
	listen = flag.String("listen", ":8022", "listen addr")
)

func mustReadByte(r *bufio.Reader) byte {
	b, err := r.ReadByte()
	if err != nil {
		panic(err)
	}

	return b
}
func readAddr(r *bufio.Reader) (add string, err error) {
	defer func() {
		e := recover() // interface{}
		if e != nil {
			err = e.(error)
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
		// log.Printf("readAddr cmd: %d, CONNECT", cmd)
	case 2:
		// log.Printf("readAddr cmd: %d, BIND", cmd)
	case 3:
		// log.Printf("readAddr cmd: %d, UDP", cmd)
	default:
		// log.Printf("readAddr cmd: %d, unknow", cmd)
		return "", errors.New("err cmd")
	}

	r.ReadByte()
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

func handshake(r *bufio.Reader, w io.Writer) (err error) {
	defer func() {
		e := recover() // interface{}
		if e != nil {
			err = e.(error)
		}
	}()
	version := mustReadByte(r)

	// log.Printf("version:%d", version)
	if version != 5 {
		return errors.New("bad Version")
	}

	nmethods, err := r.ReadByte()
	if err != nil {
		return err
	}
	// log.Printf("nmethods:%d", nmethods)

	buf := make([]byte, nmethods)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return err
	}
	// log.Printf("%v", buf)

	resp := []byte{5, 0}
	// log.Printf("resp %T\n", resp)
	_, err = w.Write(resp)
	if err != nil {
		return err
	}
	return nil
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(mycrypto.NewCryptoReader(conn, *key))
	w := mycrypto.NewCryptoWriter(conn, *key)
	err := handshake(r, w)
	if err != nil {
		log.Printf("%s handshak err %s", conn.RemoteAddr().String(), err)
		return
	}
	addr, err := readAddr(r)
	if err != nil {
		log.Printf("%s readAddr err %s", conn.RemoteAddr().String(), err)
		return
	}
	// log.Printf("addr: %s", string(addr))
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	w.Write(resp)
	server, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()
	sr := bufio.NewReader(server)
	go io.Copy(server, r)
	io.Copy(w, sr)

}

func main() {
	l, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, _ := l.Accept()
		go handleConn(conn)
	}
}
