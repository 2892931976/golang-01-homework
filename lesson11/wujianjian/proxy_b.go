package main

import (
	"flag"
	"github.com/51reboot/golang-01-homework/lesson11/wujianjian/crypto"
	"io"
	"log"
	"net"
	"sync"
)

var (
	target = flag.String("target", "www.baidu.com:80", "target host")
)

func handleConn(conn net.Conn) {
	// 建立到目标服务器的连接
	var remote net.Conn
	var err error
	remote, err = net.Dial("tcp", *target)
	if err != nil {
		log.Fatal(err)
		conn.Close()
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	// go 接收客户端的数据,发送到remote,直到conn的EOF,关闭remote
	go func() {
		defer wg.Done()
		r := NewCryptoReader(conn, "123456")
		io.Copy(remote, r) //io.Copy的连接关闭才会有返回，所以需要加go,不然会阻塞
		remote.Close()
	}()
	// go 接收remote的数据,发送到客户端,直到remote的EOF,关闭conn
	go func() {
		defer wg.Done()
		w := NewCryptoWriter(conn, "123456")
		io.Copy(w, remote)
		conn.Close()
	}()
	//等待两个协程结束
	wg.Wait()
	conn.Close()
}

func main() {
	flag.Parse()
	//建立listen
	addr := ":8021"
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
