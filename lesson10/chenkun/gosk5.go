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
	r.ReadByte()
	addrtype, _ := r.ReadByte()
	log.Printf("addr type:%d", addrtype)
	if addrtype != 3 {
		return "", errors.New("bad addr type")
	}
	// 读取一个字节的数据，代表后面紧跟的域名长度
	addrlen, _ := r.ReadByte()
	addr := make([]byte, addrlen)
	io.ReadFull(r, addr)
	log.Printf("addr1:%s", addr)

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
	fmt.Println(resp)
	conn.Write(resp)
	return nil
}

func handleconn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	handshake(r, conn) // 握手
	addr, _ := readAddr(r)
	log.Printf("addr0:%s", addr)
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)
	// 开始代理
	var remote net.Conn
	var err error
	remote, err = net.Dial("tcp", addr)
	if err != nil {
		log.Print(err)
		conn.Close()
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	// go 读取remote数据，发送到客户端conn，直到conn的EOF，关闭remote
	go func() {
		defer wg.Done()
		io.Copy(remote, conn)
		remote.Close()
	}()
	// go 读取remote数据，发送到客户端conn，直到remote的EOF，关闭conn
	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
		conn.Close()
	}()
	// 等待两个协程结束
	wg.Wait()
}

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", ":8024")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, _ := l.Accept()
		fmt.Println("1")
		fmt.Println(conn.RemoteAddr().String())
		go handleconn(conn)
	}

}
