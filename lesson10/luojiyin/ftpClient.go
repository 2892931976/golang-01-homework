package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
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

	FileName := os.Args[2]
	//_, err = conn.Write([]byte(action + " " + filepath.Base(fileName) + "\n"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	switch action {
	case "STORE":
		sendToServer(fileName, conn)
	case "GET":
		getFromServer(fileName, conn, rd)
	case "LS":
	}
	_, err = conn.Write([]byte(action + " " + filepath.Base(fileName) + "\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func writeError(conn net.Conn, err string) {
	con.Write([]byte("err: " + err))
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
