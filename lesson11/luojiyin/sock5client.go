package main

import (
	"io"
	"log"
	"net"
	"sync"
)

const key = "123456"

func main() {
	listener, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		lisConn, _ := listener.Accept()
		go handleConn(lisConn)
		log.Println("start connect to proxy server")
	}
}

func handleConn(listenConn net.Conn) error {
	log.Println("star ")
	defer listenConn.Close()

	remoteConn, err := net.Dial("tcp", "139.162.109.162:8022")
	if err != nil {
		log.Fatal(err)
	}
	defer remoteConn.Close()

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		io.Copy(remoteConn, listenConn)
	}()
	go func() {
		defer wg.Done()
		io.Copy(listenConn, remoteConn)
	}()
	wg.Wait()
	return nil
}
