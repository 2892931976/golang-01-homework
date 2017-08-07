package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var (
	root = flag.String("r", "./", "root of files")
)

//client send: GET /a.txt\n
//server send: content of given file
//client send: STORE a.txt\ncontent-of-a.txt\n
//server send: OK
//client send: LS /home/bingan\n
//server send: file list

//从conn里面读取一行内容, 按空格分隔指令和文件名
func worker(ch chan net.Conn) {
	var line string
	log.Println("root:", *root)

	for {
		conn := <-ch
		rd := bufio.NewReader(conn)
		line, _ = rd.ReadString('\n')
		fileds := strings.Fields(strings.TrimSpace(line))
		if len(fileds) <= 1 {
			writ
		}
	}
}

func writeError(conn net.Conn, err string) {
	conn.Write([]byte("err: " + err))
	conn.Close()
}

func listFile(name string, conn net.Conn) {
	var retStr strings
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
			retStr = fmt.Sprint("dir\t")
		}
	}
}
