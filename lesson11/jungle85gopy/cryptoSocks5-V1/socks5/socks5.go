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
	"sync"
)

var (
	socks5 = flag.String("s", ":8022", "port for socks5")
)

func mustReadByte(r *bufio.Reader) byte {
	b, err := r.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

// 握手。进行授权认证
func handshake(conn net.Conn, r *bufio.Reader) (err error) {
	defer func() {
		e := recover() // interface{}
		if e != nil {
			err = e.(error)
		}
	}()

	// version
	version := mustReadByte(r)
	log.Printf("version:%d", version)
	if version != 5 {
		return errors.New("bad version")
	}
	// nmethods
	nmethods := mustReadByte(r)
	log.Printf("nmethods: %d", nmethods)

	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	log.Printf("%v", buf)
	// response
	resp := []byte{5, 0}
	conn.Write(resp)
	return
}

// 获取需要代理的域名
func readAddr(r *bufio.Reader) (addrPort string, err error) {
	defer func() {
		e := recover() // interface{}
		if e != nil {
			err = e.(error)
		}
	}()

	funcName := "readAddr():"
	log.Printf("run %s ...", funcName)
	version := mustReadByte(r)
	log.Printf("%s version:%d", funcName, version)
	if version != 5 {
		return "", errors.New(funcName + " bad version")
	}
	cmd := mustReadByte(r)
	log.Printf("%s cmd:%d", funcName, cmd)
	if cmd != 1 {
		return "", errors.New(funcName + " bad cmd")
	}
	// skip rsv
	mustReadByte(r)
	// addr type
	addrType := mustReadByte(r)
	log.Printf("%s addrType:%d", funcName, addrType)
	if addrType != 3 {
		return "", errors.New(funcName + " bad addr type")
	}
	// domain name. 域名是变长的，读取一个字节的数据，代表后面紧跟着的域名的长度
	addrLen := mustReadByte(r)
	addr := make([]byte, addrLen)
	io.ReadFull(r, addr)
	log.Printf("%s domain len:%d", funcName, addrLen)
	log.Printf("%s domain:%s", funcName, addr)
	// port
	var port int16
	binary.Read(r, binary.BigEndian, &port)
	addrPort = fmt.Sprintf("%s:%d", addr, port)
	return
}

func socks5Auth(conn net.Conn, r *bufio.Reader) (addr string, err error) {
	if err = handshake(conn, r); err != nil {
		return "", err
	}
	if addr, err = readAddr(r); err != nil {
		return "", err
	}
	log.Printf("addr: %s\n", addr)
	// resp 响应客户端连接成功
	resp := []byte{0x05, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)
	return addr, nil
}

func handleConn(conn net.Conn) {
	log.Println("start handle ...")
	defer conn.Close()

	r := bufio.NewReader(conn)
	addr, err := socks5Auth(conn, r)
	if err != nil {
		log.Print(err)
		return
	}
	remote, err := net.Dial("tcp", addr)
	if err != nil {
		log.Print(err)
		conn.Close()
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		io.Copy(remote, r) // read from r, not conn
	}()
	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
	}()
	wg.Wait()
	log.Printf("shut of %s", conn.RemoteAddr().String())
	conn.Close()
	remote.Close()
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", *socks5)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("new connection from %s\n\n", conn.RemoteAddr().String())
		go handleConn(conn)
	}
}
