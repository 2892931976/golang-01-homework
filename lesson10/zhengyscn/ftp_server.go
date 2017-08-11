package main

import (
	"bufio"

	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func GetFile(conn net.Conn, r *bufio.Reader, name string) error {
	defer conn.Close()
	f, err := os.Open(name)
	if err != nil {
		log.Print(err)
		return err
	}
	defer f.Close()
	io.Copy(conn, f)
	return nil
}

func StoreFile(conn net.Conn, r *bufio.Reader, name string) error {
	// 从r读取文件内容直到err为io.EOF
	// 创建name文件
	// 向文件写入数据
	// 往conn写入OK
	// 关闭连接和文件
	defer conn.Close()
	err := os.MkdirAll(filepath.Dir(name), 0755)
	if err != nil {
		log.Printf("mkdirall error %v\n", err)
		return err
	}

	f, err := os.Create(name)
	if err != nil {
		log.Printf("create file %v\n", err)
		return err
	}
	defer f.Close()

	io.Copy(f, r)
	return nil
}

func LsDir(conn net.Conn, r *bufio.Reader, name string) error {
	err := filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		conn.Write([]byte(fmt.Sprintf("%s\n", path)))
		return nil
	})
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

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
	method := fields[0]
	name := fields[1]
	actionmap := map[string]func(net.Conn, *bufio.Reader, string) error{
		"GET":   GetFile,
		"STORE": StoreFile,
		"LS":    LsDir,
	}
	actionfunc := actionmap[method]
	if actionfunc == nil {
		log.Println("bad method")
		return
	}
	err := actionfunc(conn, r, name)
	if err != nil {
		log.Printf("execute method %s error:%s\n", method, err)
		return
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
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
