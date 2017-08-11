package main

import (
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
	action    = flag.String("a", "list", "Operation action") // 具体操作
	file_name = flag.String("n", "all", "file name")         // 操作文件名
)

func list(conn net.Conn) {
	io.Copy(os.Stdout, conn)
}

func downfile(conn net.Conn) {
	t := make([]byte, 2)
	conn.Read(t)
	if string(t) == "OK" { //下载成功的标志 写在了conn的头2个字节
		f, err := os.Create(*file_name)
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(f, conn)
	} else {
		io.Copy(os.Stdout, conn)
	}
	return
}

func upfile(conn net.Conn) {
	buf := make([]byte, 20)
	_, err := conn.Read(buf) // 接受服务端的返回
	if err != nil {
		log.Fatal("client read cmd err", err)
	}

	f, err := os.Open(*file_name) // 打开要上传的文件
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	buf, err = ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	conn.Write(buf)
	return
}
func main() {
	flag.Parse()
	addr := ":10086"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	actionfunc := map[string]func(conn net.Conn){
		"list":   list,
		"get":    downfile,
		"upload": upfile,
	}
	args := []string{*action, *file_name}

	if len(args) < 2 {
		log.Fatal("too few arguments -a list -h")
	}
	runfunc := actionfunc[*action]
	if runfunc == nil {
		log.Fatal("not found function to run")
	}
	conn.Write([]byte(strings.Join(args, " "))) // 用户输入的参数
	conn.Write([]byte("\n"))
	runfunc(conn)
}
