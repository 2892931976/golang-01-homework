package main

import (
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s file-for-put", os.Args[0])
	}
	fileName := os.Args[1]
	fd, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	addr := "127.0.0.1:8021"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err = conn.Write([]byte("STORE " + filepath.Base(fileName) + "\n"))
	if err != nil {
		log.Println("err: ", err)
	}

	num, _ := io.Copy(conn, fd)
	log.Print("write num:", num)

	// 上面io.Copy并没有主动发送结束信号
	// 接收端使用io.Copy(fd, conn)，并没有收到EOF，还在等待继续读
	// 因此使用断言，关闭发送，以便发送EOF
	value, _ := conn.(*net.TCPConn)
	value.CloseWrite()
	log.Print("after close write")

	// read
	n, err := conn.Read(buf)
	if err != nil || err == io.EOF {
		log.Print("err: ", err)
	}
	log.Printf("return content:%s", string(buf[:n]))
}
