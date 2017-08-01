package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var (
	root = flag.String("r", "./", "root of files")
	// 用于安全隔离。只能在给定的目录下访问
)

// client send: GET /a.txt\n
// server send: content of given file.
// client send: STORE a.txt\ncontent-of-a.txt\n
// server send: OK
// client send: LS /home/bingan\n
// server send: file list

// 从conn里面读取一行内容 ，按空格分隔指令和文件名
func worker(ch chan net.Conn) {
	var line string
	log.Println("root:", *root)

	for {
		conn := <-ch
		rd := bufio.NewReader(conn)
		line, _ = rd.ReadString('\n')
		fields := strings.Fields(strings.TrimSpace(line))
		if len(fields) <= 1 {
			writeError(conn, "bad input, should CMD fileName\n")
			continue
		}
		cmd, name := fields[0], fields[1]
		log.Printf("cmd:%s, name:%s\n", cmd, name)
		switch cmd {
		case "GET":
			getFile(name, conn)
		case "STORE":
			storeFile(name, conn, rd)
		case "LS":
			listFile(name, conn)
		default:
			writeError(conn, "unknown CMD\n")
			continue
		}
		conn.Close()
	}
}

func writeError(conn net.Conn, err string) {
	conn.Write([]byte("err: " + err))
	conn.Close()
}

func listFile(name string, conn net.Conn) {
	var retStr string
	fd, err := os.Open(*root + name)
	if err != nil {
		log.Print(err)
	}
	files, err := fd.Readdir(-1)
	if err != nil {
		log.Print(err)
	}
	conn.Write([]byte("type\tname\t\tsize\n"))
	for _, f := range files {
		if f.IsDir() {
			retStr = fmt.Sprintf("dir\t%s\t%d\n", f.Name(), f.Size())
		} else {
			retStr = fmt.Sprintf("file\t%s\t%d\n", f.Name(), f.Size())
		}
		conn.Write([]byte(retStr))
	}
	conn.Close()
}

func storeFile(name string, conn net.Conn, rd *bufio.Reader) {
	fd, err := os.Create(*root + name)
	if err != nil {
		log.Print(err)
		return
	}
	defer fd.Close()

	// 注意读写文件与socket时，io.Copy从bufio和conn读取的差别
	n, err := io.Copy(fd, rd)
	log.Print("read num:", n)
	conn.Write([]byte("OK"))
	conn.Close()
}

// 打开文件，读取内容，发送内容，关闭连接和文件
func getFile(name string, conn net.Conn) {
	fd, err := os.Open(*root + name)
	if err != nil {
		conn.Write([]byte(err.Error()))
		// 不能使用Fatal，否则服务端会因此退出
		log.Println(err)
	}
	defer fd.Close()
	io.Copy(conn, fd)
}

func server(listener net.Listener) {
	connCh := make(chan net.Conn)
	go worker(connCh)
	go worker(connCh)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		connCh <- conn
	}
}

func main() {
	flag.Parse()

	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	server(listener)
}
