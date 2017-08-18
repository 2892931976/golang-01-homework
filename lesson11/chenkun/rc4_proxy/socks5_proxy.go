package main

import (
	"bufio"
	"encoding/binary"
	"errors"
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

func handshake(r *bufio.Reader, clientconn net.Conn) error {
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
	clientconn.Write(resp)
	return nil
}

func handleconn(clientconn net.Conn) {
	defer clientconn.Close()
	r := bufio.NewReader(clientconn)
	handshake(r, clientconn) // 握手
	addr, _ := readAddr(r)
	log.Printf("addr0:%s", addr)
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	clientconn.Write(resp)
	// 开始代理
	var serverconn net.Conn
	var err error
	serverconn, err = net.Dial("tcp", addr)
	if err != nil {
		log.Print(err)
		clientconn.Close()
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	// go 读取serverconn数据，发送到客户端clientconn，直到clientconn的EOF，关闭serverconn
	go func() {
		defer wg.Done()
		io.Copy(serverconn, clientconn)
		serverconn.Close()
	}()
	// go 读取serverconn数据，发送到客户端clientconn，直到serverconn的EOF，关闭clientconn
	go func() {
		defer wg.Done()
		io.Copy(clientconn, serverconn)
		clientconn.Close()
	}()
	// 等待两个协程结束
	wg.Wait()
}

func main() {
	l, err := net.Listen("tcp", ":8857")
	if err != nil {
		log.Fatal(err)
	}
	for {
		clientconn, _ := l.Accept()
		go handleconn(clientconn)
	}

}
