package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

var (
	root = flag.String("root", "/", "root of ftp server data dir")
)

// client -> GET /home/binggan/a.txt\n
// server -> content of /home/bingan/a.txt

// client -> STORE /home/bingan/a.txt\n content of file EOF
// server -> OK

// client -> LS /home/bingan\n
// server -> content of dir /home/bingan

func handleConn(conn net.Conn) {
	// 从conn里面读取一行内容
	// 按空格分割指令和文件名
	// 打开文件
	// 读取内容
	// 发送内容
	// 关闭连接和文件
	defer conn.Close()
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len(fields) != 2 {
		conn.Write([]byte("bad input"))
		return
	}
	cmd := fields[0]
	name := fields[1]
	if cmd == "GET" {
		f, err := os.Open(name)
		if err != nil {
			log.Print(err)
			return
		}
		defer f.Close()
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			log.Print(err)
			return
		}
		conn.Write(buf)
	} else if cmd == "STORE" {
		// 从r读取文件内容直到err为io.EOF
		// 创建name文件
		// 向文件写入数据
		// 往conn写入OK
		// 关闭连接和文件
	}
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
