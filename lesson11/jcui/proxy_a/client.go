package main

import (
	"github.com/51reboot/golang-01-homework/lesson10/jcui/mycrypto"
	"io"
	"log"
	"net"
	"sync"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	//这里的9999 指请求远端sevrer的端口
	remote, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Print(err)
		return
	}
	defer remote.Close()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	key := "AB234asfds345safdasd"
	// 将conn请求的数据加密后 通过io.Copy复制给remote
	go func() {
		defer wg.Done()
		w := mycrypto.NewCryptoWriter(remote, key)
		io.Copy(w, conn)
	}()

	// 将conn请求的数据加密后 通过io.Copy复制给remote
	go func() {
		defer wg.Done()
		w := mycrypto.NewCryptoReader(remote, key)
		io.Copy(conn, w)
	}()
	wg.Wait()
}

func main() {
	//这里7777 指客户端启动监听的端口
	listen, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	for {
		conn, _ := listen.Accept()
		go handleConn(conn)
		log.Print(conn.RemoteAddr().String())
	}
}
