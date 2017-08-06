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
	port = flag.String("port", ":8022", "port for proxy")
)

// 握手。进行授权认证
func handshake(r *bufio.Reader, conn net.Conn) error {
	// version
	version, _ := r.ReadByte()
	log.Printf("version:%d", version)
	if version != 5 {
		return errors.New("bad version")
	}
	// nmethods
	nmethods, _ := r.ReadByte()
	log.Printf("nmethods: %d", nmethods)
	// methods
	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	log.Printf("%v", buf)
	// response
	resp := []byte{5, 0}
	conn.Write(resp)
	return nil
}

// 获取需要代理的域名
func readAddr(r *bufio.Reader) (string, error) {
	funcName := "readAddr():"
	log.Printf("run %s ...", funcName)
	version, err := r.ReadByte()
	log.Printf("%s version:%d", funcName, version)
	if err != nil {
		return "", err
	}
	if version != 5 {
		return "", errors.New(funcName + " bad version")
	}
	cmd, err := r.ReadByte()
	log.Printf("%s cmd:%d", funcName, cmd)
	if err != nil {
		return "", err
	}
	if cmd != 1 {
		return "", errors.New(funcName + " bad cmd")
	}
	// skip rsv
	r.ReadByte()
	// addr type
	addrType, _ := r.ReadByte()
	log.Printf("%s addrType:%d", funcName, addrType)
	if addrType != 3 {
		return "", errors.New(funcName + " bad addr type")
	}
	// domain name. 域名是变长的，读取一个字节的数据，代表后面紧跟着的域名的长度
	addrLen, _ := r.ReadByte()
	addr := make([]byte, addrLen)
	io.ReadFull(r, addr)
	log.Printf("%s domain len:%d", funcName, addrLen)
	log.Printf("%s domain:%s", funcName, addr)
	// port
	var port int16
	binary.Read(r, binary.BigEndian, &port)
	return fmt.Sprintf("%s:%d", addr, port), nil
}

func socks5Auth(conn net.Conn) (string, error) {
	r := bufio.NewReader(conn)
	if err := handshake(r, conn); err != nil {
		return "", err
	}
	addr, err := readAddr(r)
	if err != nil {
		return "", err
	}
	log.Printf("addr: %s\n", addr)
	// resp 响应客户端连接成功
	resp := []byte{5, 0, 0, 3, 0, 0, 0, 0, 0, 0}
	conn.Write(resp)
	return addr, nil
}

func handleConn(conn net.Conn) {
	log.Println("start handle ...")
	defer conn.Close()

	addr, err := socks5Auth(conn)
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
		io.Copy(remote, conn)
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
	listener, err := net.Listen("tcp", *port)
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
