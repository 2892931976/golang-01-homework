package main

import (
	"flag"
	"fmt"
	"io"
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
	f, err := os.Create(*file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//buf := make([]byte, 4096)
	//for {
	//	n, err :=conn.Read(buf)
	//	if err == io.EOF{
	//		break
	//	}
	//	f.Write(buf[:n])
	//
	//}

	io.Copy(f, conn)

	return
}

func upfile(conn net.Conn) {

	f, err := os.Open(*file_name) // 打开要上传的文件
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	//buf := make([]byte, 4096)
	//for {
	//	n, err :=f.Read(buf)
	//	if err != nil || err == io.EOF{
	//		break
	//	}
	//	conn.Write(buf[:n])
	//}
	io.Copy(conn, f)
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
