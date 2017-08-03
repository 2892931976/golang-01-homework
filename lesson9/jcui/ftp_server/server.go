package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

var (
	root = flag.String("r", "./", "ftp root director")
	// 用于定义一个指定目录,便于ftp的安全
)

/*
监听端口
接受新的连接
启动协程
发送接收数据
断开连接
*/

func Getfile(name string, conn net.Conn) {
	r, err := os.Open(*root + name)
	if err != nil {
		log.Print(err)
		return
	}
	defer r.Close()
	//读取内容
	buf, err := ioutil.ReadAll(r)
	////发送内容
	conn.Write(buf)
}

func Listfile(name string, conn net.Conn) {
	r, err := os.Open(*root + name)
	if err != nil {
		log.Print(err)
		return
	}
	defer r.Close()
	dirs, err := r.Readdir(-1)
	if err != nil {
		log.Print(err)
		return
	}
	for _, f := range dirs {
		if f.IsDir() {
			conn.Write([]byte(fmt.Sprintf("dir\t%s\t%d\n", f.Name(), f.Size())))
		} else {
			conn.Write([]byte(fmt.Sprintf("file\t%s\t%d\n", f.Name(), f.Size())))
		}
	}
}

func Storefile(name string, conn net.Conn, r *bufio.Reader) {
	f, err := os.Create(*root + name)
	if err != nil {
		log.Print(err)
		return
	}
	defer f.Close()

	// 注意读写文件与socket时，io.Copy从bufio和conn读取的差别
	io.Copy(f, r)
	conn.Write([]byte("Store files OK \n"))
}

func handleConn(conn net.Conn) { //主机conn 这里的类型net.Conn
	defer conn.Close()
	//读取客户端需求,获得客户端需要得到的文件
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len(fields) != 2 {
		conn.Write([]byte("bad input \n"))
		return
	}
	cmd := fields[0]
	name := fields[1]
	log.Printf("%s %s %s\n", conn.RemoteAddr().String(), cmd, name)
	switch cmd {
	case "GET":
		Getfile(name, conn)
	case "LS":
		Listfile(name, conn)
	case "STORE":
		Storefile(name, conn, r)
		// 从r读取文件直到err为io.EOF
		//创建name文件
		//向文件写入数据
		//往conn写入OK
		//关闭连接和文件
	}
	defer conn.Close()
}

func main() {
	addr := ":7777" //监听任意IP的7777端口
	//创建监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	//
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)

	}
}
