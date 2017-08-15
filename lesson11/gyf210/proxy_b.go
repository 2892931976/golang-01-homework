package main

import (
	"flag"
	"github.com/gyf210/mycrypto"
	"io"
	"log"
	"net"
	"sync"
)

var (
	target = flag.String("t", "127.0.0.1:9005", "target addr")
	listen = flag.String("l", ":9001", "listen addr")
	key    = flag.String("k", "123456", "crypto key")
)

func handle(conn net.Conn) {
	defer conn.Close()
	remote, err := net.Dial("tcp", *target)
	if err != nil {
		log.Println(err)
		return
	}
	defer remote.Close()
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		r := mycrypto.NewCryptoReader(conn, *key)
		io.Copy(remote, r)
		remote.Close()
	}()
	go func() {
		defer wg.Done()
		w := mycrypto.NewCryptoWriter(conn, *key)
		io.Copy(w, remote)
		conn.Close()
	}()
	wg.Wait()
}

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
