package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func get(conn net.Conn, fname string) {
	f, err := os.Open(fname)
	if err != nil || err == io.EOF {
		log.Print(err)
		return
	}
	defer f.Close()
	io.Copy(conn, f)
}

func store(r io.Reader, conn net.Conn, fname string) {
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Print(err)
		conn.Write([]byte("fail."))
		return
	}
	defer f.Close()
	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			log.Print(err)
			break
		}
		f.Write(buf[:n])
	}
	conn.Write([]byte("OK"))
}

func ls(conn net.Conn, path string) {
	finfo, err := os.Stat(path)
	if err != nil {
		conn.Write([]byte("fail\n"))
		return
	}
	mode := finfo.Mode()
	if mode.IsRegular() {
		result := fmt.Sprintf("%s is not dirctory.\n", path)
		conn.Write([]byte(result))
		return
	}
	fs, err := ioutil.ReadDir(path)

	if err != nil {
		result := fmt.Sprintf("%s can not open.\n", path)
		conn.Write([]byte([]byte(result)))
		return
	}
	for _, f := range fs {
		result := fmt.Sprintf("%s\t", f.Name())
		conn.Write([]byte(result))
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len(fields) != 2 {
		conn.Write([]byte("bad input\n"))
		return
	}
	cmd := fields[0]
	name := fields[1]

	switch cmd {
	case "GET":
		get(conn, name)
	case "STORE":
		store(r, conn, name)
	case "LS":
		ls(conn, name)
	default:
		conn.Write([]byte("hello, golang\n"))
	}
	conn.Write([]byte("\n"))
}
func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
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
