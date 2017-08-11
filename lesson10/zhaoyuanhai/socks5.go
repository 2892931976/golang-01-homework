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

//1. 握手
//2. 获取客户端的请求
//3. 开始代理
//
//
func handshake(r *bufio.Reader, conn net.Conn) error {
	//获取客户端协议版本，1个字节，ReadByte表示读取1个字节
	version, _ := r.ReadByte()
	log.Printf("version:%d\n", version)
	if version != 5 {
		return errors.New("bad version")
	}
	//读取认证方式长度，1个字节
	nmethods, _ := r.ReadByte()
	log.Printf("nmethods:%d\n", nmethods)

	//读取认证方式，根据认证方式长度，创建对应的BUF，当BUF读满则完整读出认证方式
	buf := make([]byte, nmethods)
	io.ReadFull(r, buf) //从r中读取内容，直到buf填充满
	log.Printf("buf: %v", buf)

	//给客户端返回服务端支持的认证方式及长度
	resp := []byte{5, 0}
	conn.Write(resp)
	return nil
}

func readAddr(r *bufio.Reader) (string, error) {
	//读取客户端发来的请求协议信息，1个字节
	version, _ := r.ReadByte()
	log.Printf("Version: %d\n", version)
	if version != 5 {
		return "", errors.New("bad version")
	}

	//读取客户端发来的代理请求类型，1个字节
	cmd, _ := r.ReadByte()
	log.Printf("client request type: %d\n", cmd)
	if cmd != 1 {
		return "", errors.New("bad client type")
	}

	//读取保留字，1个字节，跳过
	r.ReadByte()

	//读取addr type字段，1个字节
	atyp, _ := r.ReadByte()
	log.Printf("%v", atyp)

	if atyp != 3 {
		return "", errors.New("bad address type")
	}

	// rt := fmt.Sprintf("%v--%v--%v", version, cmd, atyp)

	//读取一个字节的数据，代表后面紧跟着的域名长度
	addlen, _ := r.ReadByte()
	log.Printf("addr len: %d\n", addlen)

	//读取n个字节得到域名，n根据上一步得到的结果来决定
	addr := make([]byte, addlen)
	io.ReadFull(r, addr)
	log.Printf("addr buf: %s\n", addr)

	//读取端口
	//参数一，参数二：BigEndian大断续，参数三：数据存放的变量地址
	var port int16
	binary.Read(r, binary.BigEndian, &port)
	log.Printf("addr port: %d\n", port)

	return fmt.Sprintf("%s:%d", addr, port), nil
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	handshake(r, conn)
	addr, _ := readAddr(r)
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)
	proxy(conn, addr)
}

func proxy(conn net.Conn, address string) {

	toDst, err := net.Dial("tcp", address)
	if err != nil {
		log.Print(err)
		conn.Close()
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		io.Copy(toDst, conn)
		toDst.Close()
	}()

	//go 从目标服务器发送到客户端的协程
	go func() {
		defer wg.Done()
		io.Copy(conn, toDst)
		conn.Close()
	}()
	//等待两个协程结束
	wg.Wait()
}

func main() {
	// flag.Parse()
	// //建立listen
	// addr := ":7777"
	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		//接受新的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}
