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

	"github.com/jungle85gopy/learn-go/51reboot/cryptoSocks5-V1/mycrypto"
)

var (
	socks5 = flag.String("s", ":8021", "port for socks5")
)

const key = "123456"

func mustReadByte(r *bufio.Reader) byte {
	b, err := r.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

// 握手。进行授权认证
func handshake(r *bufio.Reader, wr io.Writer) (err error) {
	defer func() {
		e := recover() // interface{}
		if e != nil {
			err = e.(error)
		}
	}()

	// version
	version := mustReadByte(r)
	if version != 5 {
		return errors.New("bad version")
	}
	// nmethods
	nmethods := mustReadByte(r)
	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	// response
	resp := []byte{5, 0}
	_, err = wr.Write(resp[:])
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
	version := mustReadByte(r)
	if version != 5 {
		return "", errors.New(funcName + " bad version")
	}
	cmd := mustReadByte(r)
	if cmd != 1 {
		return "", errors.New(funcName + " bad cmd")
	}
	// skip rsv
	mustReadByte(r)
	// addr type
	addrType := mustReadByte(r)
	if addrType != 3 {
		return "", errors.New(funcName + " bad addr type")
	}
	// domain name. 域名是变长的，读取一个字节的数据，代表后面紧跟着的域名的长度
	addrLen := mustReadByte(r)
	addr := make([]byte, addrLen)
	io.ReadFull(r, addr)
	log.Printf("%s domain:%s", funcName, addr)
	// port
	var port int16
	binary.Read(r, binary.BigEndian, &port)
	addrPort = fmt.Sprintf("%s:%d", addr, port)
	return
}

func socks5Auth(rd *bufio.Reader, wr io.Writer) (addr string, err error) {
	if err = handshake(rd, wr); err != nil {
		return "", err
	}
	if addr, err = readAddr(rd); err != nil {
		return "", err
	}
	// resp 响应客户端连接成功
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	_, err = wr.Write(resp)
	return
}

func handleConn(conn net.Conn) {
	log.Println("start handle ...")
	defer conn.Close()

	// 封装client连接
	clientRd := mycrypto.NewCryptoReader(conn, key)
	clientWr := mycrypto.NewCryptoWriter(conn, key)

	rdBuf := bufio.NewReader(clientRd)
	addr, err := socks5Auth(rdBuf, clientWr)
	if err != nil {
		log.Printf("get addr err:%s\n", err.Error())
		return
	}
	remote, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("dial remote err:%s", err.Error())
		conn.Close()
		return
	}
	defer remote.Close()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	// from (conn -> NewCryptoReader -> bufio.NewReader) to remote
	go func() {
		defer wg.Done()
		io.Copy(remote, rdBuf)
	}()
	// from remote to client writer
	go func() {
		defer wg.Done()
		io.Copy(clientWr, remote)
	}()
	wg.Wait()
	log.Printf("shut of %s", conn.RemoteAddr().String())
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
