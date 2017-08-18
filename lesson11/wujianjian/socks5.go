package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

// 1. 握手
// 2. 获取客户端代理请求
// 3. 开始代理

func readAddr(r *bufio.Reader) (string, error) {
	version, _ := r.ReadByte()
	log.Printf("version:%d", version)
	if version != 5 {
		return "", errors.New("bad version")
	}
	cmd, _ := r.ReadByte()
	log.Printf("cmd:%d", cmd)
	if cmd != 1 {
		return "", errors.New("bad cmd")
	}

	// skip rsv字段
	r.ReadByte()

	addrtype, _ := r.ReadByte()
	log.Printf("addr type:%d", addrtype)
	if addrtype != 3 {
		return "", errors.New("bad addr type")
	}

	// 读取一个字节的数据，代表后面紧跟着的域名长度
	// 读取n个字节得到域名，n根据上一步得到的结果来决定
	// addrlen
	// addr
	addrlen, _ := r.ReadByte()
	log.Printf("addrlen:%d", addrlen)

	addr := make([]byte, addrlen)
	io.ReadFull(r, addr) //从reader读取直到把buf填充满
	log.Printf("%s", addr)

	//读取2个字节的方式
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
	io.ReadFull(r, buf) //从reader读取直到把buf填充满
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
	fmt.Printf("addr:%s", addr)
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)
	// 3.开始代理
	// 建立到目标服务器的连接
	var remote net.Conn
	var err error
	remote, err = net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
		conn.Close()
		return
	}
	defer remote.Close()
	go io.Copy(remote, r)
	io.Copy(conn, remote)

}

func main() {
	//建立listen
	addr := ":8022"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		// accept new connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
