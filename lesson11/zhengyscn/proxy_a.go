package main

import (
	"flag"
	"github.com/467754239/godoc/lesson/lesson11/mycrypto"
	"io"
	"log"
	"net"
	"sync"
)

var (
	target = flag.String("target", "www.baidu.com:80", "target host")
)

func handleConn(conn net.Conn) {
	remote_conn, err := net.Dial("tcp", *target)
	log.Printf("target:%s", *target)
	if err != nil {
		log.Print(err)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		w := mycrypto.NewCryptoWriter(remote_conn, "123456")
		io.Copy(w, conn) // Connection closed or EOF
		defer remote_conn.Close()
	}()
	go func() {
		defer wg.Done()
		r := mycrypto.NewCryptoReader(remote_conn, "123456")
		io.Copy(conn, r)
		defer conn.Close()
	}()
	wg.Wait()

}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}
