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

func handshake(r *bufio.Reader, conn net.Conn) error {
	version, _ := r.ReadByte() //ReadByte 代表读取一个字节
	//处理第一个字节
	log.Printf("version:%d", version)
	if version != 5 {
		return errors.New("bad version")
	}
	//处理第二个字节
	nmethods, _ := r.ReadByte()
	buf := make([]byte, nmethods)
	io.ReadFull(r, buf) //将buf填充满
	log.Printf("%v", buf)

	//返回数据
	resp := []byte{5, 0}
	conn.Write(resp)
	return nil
}

func readAddr(r *bufio.Reader) (string, error) {
	//处理第一个数据
	version, _ := r.ReadByte()
	log.Printf("version:%d", version)
	if version != 5 {
		return "", errors.New("bad version")
	}

	//处理第二个数据
	cmd, err := r.ReadByte()
	if err != nil {
		return "", err
	}
	log.Printf("%s", cmd)
	if cmd != 1 {
		return "", errors.New("bad cmd")
	}

	//处理第三个数据(保留数据跳过即可)
	r.ReadByte()

	//处理第四个数据
	readtype, _ := r.ReadByte()
	log.Printf("%s", readtype)
	if readtype != 3 {
		return "", errors.New("bad type")
	}

	//处理第五个数据
	/*
		读取一个字节的数据,代表后面竟跟着域名的长度
		读取n个字节的得到域名,n根据上一步得到的结果来决定
		addrlen  长度
		addr     地址
	*/
	addrlen, _ := r.ReadByte()
	addr := make([]byte, addrlen)
	io.ReadFull(r, addr)
	log.Printf("addrlen:%s", addrlen)
	log.Printf("addr:%s", addr)

	//处理第六个数据,占位2个字节
	var port int16
	binary.Read(r, binary.BigEndian, &port)

	return fmt.Sprintf("%s:%d", addr, port), nil
}

func handleConn(conn net.Conn) {
	/*
		参考链接地址: http://www.jianshu.com/p/172810a70fad

		1.握手
		socks5 接收3个字段的数据
			第一个字段 1个字节  VER 代表版本
			第二个字段 1个字节  NMETHODS 表示第三个字段的长度
			第三个字段 1-255个字节 METHODS 表示客户端支持的认证方式
		2.获取客户端代理的请求
			第一个字段 1个字节  VER 代表版本
			第二个字段 1个字节
			第三个字段 1个字节
			第四个字段 1个字节
			第五个字段
			第六个字段 2个字节
		3.开始代理
	*/
	r := bufio.NewReader(conn)
	//开始代理
	handshake(r, conn)
	//获取客户端代理的请求
	addr, err := readAddr(r)
	if err != nil {
		log.Print("Addr Error:", err)
	}
	log.Printf("addr:", addr)
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)
	//开始代理
	var remote net.Conn

	remote, err = net.Dial("tcp", addr)
	if err != nil {
		log.Print("Error:", err)
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	//go 接收客户端的数据,发送到remote,直到conn的EOF
	go func() {
		defer wg.Done()
		io.Copy(remote, conn)
		remote.Close()
	}()
	//go 接收remote的数据,发送到客户端,直到remote的EOF
	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
		conn.Close()
	}()
	//等待两个协程结束
	wg.Wait()
}

func main() {
	//建立监听
	addr := ":9000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	//接收新的连接
	for {
		//accept new conection
		conn, err := listener.Accept()
		log.Print(conn.RemoteAddr())
		if err != nil {
			log.Print(err)
		}
		// 参考页面 http://www.jianshu.com/p/172810a70fad
		go handleConn(conn)
	}
}
