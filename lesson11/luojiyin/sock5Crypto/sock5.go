package main

import (
	"bufio"
	"crypto/md5"
	"crypto/rc4"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const key = "123456"

func readAddr(r *bufio.Reader) (string, error) {
	version, _ := r.ReadByte()
	log.Printf("readAddr version: %d", version)
	if version != 5 {
		return "", errors.New("bad version")
	}
	cmd, _ := r.ReadByte()
	switch cmd {
	case 1:
		log.Printf("readAddr cmd: %d CONNECT", cmd)
	case 2:
		log.Printf("readAddr cmd: %d BIND", cmd)
	case 3:
		log.Printf("readAddr cmd: %d UDP", cmd)
	default:
		log.Printf("readAddr cmd %d , unknow", cmd)
		return "", errors.New("err cmd")
	}

	r.ReadByte() //跳个字节，此字节无意义
	atyp, _ := r.ReadByte()
	var addrlen int
	switch atyp {
	case 1:
		log.Printf("readAddr atyp: %d, ipv4", atyp)
		addrlen = 4
	case 3:
		log.Printf("readAddr atyp: %d, domin", atyp)
		addrl, _ := r.ReadByte()
		log.Println(addrl)
		addrlen = int(addrl)
	case 4:
		log.Printf("readAddr atyp: %d, ipv6", atyp)
		addrlen = 16
	}
	host := make([]byte, addrlen)
	var hh string
	io.ReadFull(r, host)
	log.Printf("host:%s\n", host)
	switch atyp {
	case 1:
		hh = net.IPv4(host[0], host[1], host[2], host[3]).String()
	case 3:
		hh = string(host)
		log.Printf("hh:%s", hh)
	}
	log.Printf("readAddr domain: %s ", hh)

	var port uint16
	binary.Read(r, binary.BigEndian, &port)

	return fmt.Sprintf("%s:%d", hh, port), nil
}

func handshake(r *bufio.Reader, conn net.Conn) error {
	version, _ := r.ReadByte()
	log.Printf("version: %d ", version)
	if version != 5 {
		return errors.New("bad Version")
	}
	nmethods, _ := r.ReadByte()
	log.Printf("nmethods: %d", nmethods)

	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	log.Printf("%v", buf)

	resp := []byte{5, 0}
	log.Printf("resp %T\n", resp)
	conn.Write(resp)
	return nil
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	handshake(r, conn)
	addr, _ := readAddr(r)
	log.Printf("addr: %s", string(addr))
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)
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
	l, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, _ := l.Accept()
		go handleConn(conn)
	}
}

type CryptoWriter struct {
	w      io.Writer
	cipher *rc4.Cipher
}

func NewCrytoWrite(w io.Writer, key string) io.Writer {
	md5Sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5Sum[:])
	if err != nil {
		log.Print(err)
	}
	return &CryptoWriter{
		w:      w,
		cipher: cipher,
	}
}

func (w *CryptoWriter) Write(b []byte) (int, error) {
	buf := make([]byte, len(b))
	w.cipher.XORKeyStream(buf, b)
	w.w.Write(buf)
	return len(buf), nil
}

type CryptoReader struct {
	r      io.Reader
	cipher *rc4.Cipher
}

func NewCrytoReader(r io.Reader, key string) io.Reader {
	md5Sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5Sum[:])
	if err != nil {
		log.Print(err)
	}
	return &CryptoReader{
		r:      r,
		cipher: cipher,
	}
}

func (r *CryptoReader) read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	buf := b[:n]
	r.cipher.XORKeyStream(buf, buf)
	return n, err
}
