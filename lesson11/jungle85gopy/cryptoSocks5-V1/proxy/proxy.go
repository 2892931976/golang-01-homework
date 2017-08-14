package main

import (
	"flag"
	"io"
	"log"
	"net"
	"sync"

	"github.com/jungle85gopy/learn-go/51reboot/cryptoSocks5-V1/mycrypto"
)

var (
	client = flag.String("c", ":8021", "client port")
	proxy  = flag.String("p", ":8022", " proxy port")
)

const key = "123456"

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", *client)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		lisConn, _ := listener.Accept()
		go handleConn(lisConn)
		log.Printf("new connection from %s\n\n", lisConn.RemoteAddr().String())
	}
}

func handleConn(listenConn net.Conn) {
	log.Println("start handle ...")
	defer listenConn.Close()

	remoteConn, err := net.Dial("tcp", *proxy)
	if err != nil {
		log.Fatal(err)
	}
	defer remoteConn.Close()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	// 对conn进行解密Reader封装，解密后Copy到remote，直到conn端EOF
	go func() {
		defer wg.Done()
		lRd := mycrypto.NewCryptoReader(listenConn, key)
		io.Copy(remoteConn, lRd)
	}()
	// 对conn进行加密Writer封装，remote数据Copy到Writer，直到remote端EOF
	go func() {
		defer wg.Done()
		rWr := mycrypto.NewCryptoWriter(listenConn, key)
		io.Copy(rWr, remoteConn)
	}()
	wg.Wait()
	log.Printf("shut of listen of %s", listenConn.RemoteAddr().String())
}
