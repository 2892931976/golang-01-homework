package main

import (
	"io"
	"log"
	"net"
	"rc4_proxy/encrypt_decrypt"
	"sync"
)

var (
	target = "127.0.0.1:8856" // rc4_proxy_b的地址端口
)

func handleconn(clientconn net.Conn) {
	key := "123456"
	var serverconn net.Conn
	var err error
	serverconn, err = net.Dial("tcp", target)
	if err != nil {
		log.Print(err)
		clientconn.Close()
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	// go 读取clientconn数据，加密，发送到serverconn，直到clientconn的EOF，关闭serverconn
	go func() {
		defer wg.Done()
		io.Copy(cpt.NewCryptoWriter(serverconn, key), clientconn)
		serverconn.Close()
	}()
	// go 读取serverconn数据，解密，发送到clientconn，直到serverconn的EOF，关闭clientconn
	go func() {
		defer wg.Done()
		io.Copy(clientconn, cpt.NewCryptoReader(serverconn, key))
		clientconn.Close()
	}()
	// 等待两个协程结束
	wg.Wait()
}

func main() {
	listener, err := net.Listen("tcp", ":8855")
	if err != nil {
		log.Print(err)
	}
	defer listener.Close()

	for {
		clientconn, _ := listener.Accept()
		go handleconn(clientconn)
	}
}
