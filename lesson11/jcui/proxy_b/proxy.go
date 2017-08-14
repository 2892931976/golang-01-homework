package main

import (
	"golang-01/lesson10/jcui/mycrypto"
	"io"
	"log"
	"net"
	"sync"
)

func handleConn_accept(conn net.Conn) {
	defer conn.Close()
	//将数据解密
	key := "AB234asfds345safdasd"
	remote, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Print(err)
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		r := mycrypto.NewCryptoReader(conn, key)
		io.Copy(remote, r)
	}()
	go func() {
		defer wg.Done()
		w := mycrypto.NewCryptoWriter(conn, key)
		io.Copy(w, remote)
	}()
	wg.Wait()
}

func main() {
	listen_accept, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}

	defer listen_accept.Close()
	for {
		conn_s, err1 := listen_accept.Accept()
		if err1 != nil {
			log.Print(err1)
		}
		go handleConn_accept(conn_s)
		log.Print("s:", conn_s.RemoteAddr().String())
	}
}
