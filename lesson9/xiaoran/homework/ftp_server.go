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
	"path/filepath"
	"strings"
)

var (
	root = flag.String("root", "/", "root of ftp server data dir")
)

//client -> GET /home/bingan/a.txt\n
//server -> content of /home/bingan/a.txt\n

//client -> STORE /home/bingan/a.txt\n content of file EOF
//server -> OK

//client -> LS /home/bingan\n
//server -> content of dir /home/bingan

func get(args []string, conn net.Conn) error {
	filename := args[0]
	f, err := os.Open(filename)
	if err != nil {
		log.Print(err)
		//return
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Print(err)
		//return
	}
	conn.Write(buf)
	return nil
}

func store(args []string, conn net.Conn) error {
	//从r读取文件内容直到err为io.EOF
	//创建name文件
	//向文件写入数据
	//往conn写入OK
	//关闭连接和文件
	filename := args[0]
	f, err := os.Open(filename)
	if err != nil {
		log.Print(err)
		//return
	}
	defer f.Close()

	name := filepath.Base(filename)
	ftp_file, err := os.Create(name)
	if err != nil {
		log.Print(err)
	}
	defer ftp_file.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print("line:", line)
		io.Copy(ftp_file, r)
	}
	conn.Write([]byte("OK"))
	return nil
}

func handleConn(conn net.Conn) {
	//从conn里面读取一行内容
	//按空格分割指令和文件名
	//打开文件
	//读取内容
	//发送内容
	//关闭连接和文件
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
	args := fields[1:]

	actionmap := map[string]func([]string, net.Conn) error{
		"GET":   get,
		"STORE": store,
	}
	actionfunc := actionmap[cmd]
	if actionfunc == nil {
		fmt.Println("输入命令有误")
	}

	err := actionfunc(args, conn)
	if err != nil {
		fmt.Printf("execute action")
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
