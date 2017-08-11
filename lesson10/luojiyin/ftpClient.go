package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

var (
	root = flag.String("r", "./", "root of files")
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("LS can kown files in server, GET fileName get file info\n  STORE file can upload file")
	}

	addr := "127.0.0.1:8021"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	rd := bufio.NewReader(conn)
	defer conn.Close()

	buf := make([]byte, 1024)

	action := os.Args[1]

	fileName := os.Args[2]
	_, err = conn.Write([]byte(action + " " + filepath.Base(fileName) + "\n"))
	if err != nil {
		log.Fatal(err)
	}
	switch action {
	case "STORE":
		//sendToServer(fileName, conn)
		fd, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()
		defer conn.Close()

		//buf := make([]byte, 1024)
		num, _ := io.Copy(conn, fd)
		log.Print("write num:", num)

		value, _ := conn.(*net.TCPConn)
		value.CloseWrite()
		log.Print("after close write")

		n, err := conn.Read(buf)
		if err != nil || err == io.EOF {
			log.Print("err :", err)
		}
		log.Printf("return content:%s", string(buf[:n]))
	case "GET":
		//getFromServer(fileName, conn, rd)
		fd, err := os.Create(*root + fileName)
		if err != nil {
			log.Print(err)
			return
		}
		defer fd.Close()
		n, err := io.Copy(fd, rd)
		log.Print("read num :", n)
		conn.Write([]byte("ok"))
		conn.Close()

	case "LS":
		n, err := conn.Read(buf)
		if err != nil || err == io.EOF {
			log.Print("err : ", err)
		}
		log.Printf("return content :%s", string(buf[:n]))

	}
}

func writeError(conn net.Conn, err string) {
	conn.Write([]byte("err: " + err))
	conn.Close()
}

func sendToServer(name string, conn net.Conn) {
	fd, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	io.Copy(conn, fd)

}

func getFromServer(name string, conn net.Conn, rd *bufio.Reader) {
	fd, err := os.Create(name)
	if err != nil {
		log.Print(err)
		return
	}
	defer fd.Close()

	n, err := io.Copy(fd, rd)
	log.Print("get num: ", n)
	conn.Write([]byte("ok"))
	conn.Close()
}
