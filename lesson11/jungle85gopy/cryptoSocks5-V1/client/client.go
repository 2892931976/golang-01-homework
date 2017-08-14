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
	client = flag.String("c", ":8020", "client port")
	proxy  = flag.String("p", ":8021", " proxy port")
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

	// 读取conn数据，将其Copy到加密Writer，加密后写到remote，直到conn端EOF
	go func() {
		defer wg.Done()
		rWr := mycrypto.NewCryptoWriter(remoteConn, key)
		io.Copy(rWr, listenConn)
	}()
	// 读取remote的数据，并经过解密，然后发送给客户端conn，直到远端EOF
	go func() {
		defer wg.Done()
		lRd := mycrypto.NewCryptoReader(remoteConn, key)
		io.Copy(listenConn, lRd)
	}()
	wg.Wait()
	log.Printf("shut of listen of %s", listenConn.RemoteAddr().String())
}
