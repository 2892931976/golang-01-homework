package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var (
	root = flag.String("r", "./", "ftp root director")
	// 用于定义一个指定目录,便于ftp的安全
)

func handConn(conn net.Conn) { //conn类型为net.conn
	//从conn读取一行内容
	//按空格分隔指令和文件名
	defer conn.Close() //函数退出，关闭连接
	r := bufio.NewReader(conn)
	content, err := r.ReadString('\n')
	if err != nil {
		log.Print(err)
		return
	}
	content = strings.TrimSpace(content)
	cmdfn := strings.Fields(content)
	if len(cmdfn) != 2 {
		conn.Write([]byte("bad input"))
		return
	}
	fmt.Println(cmdfn[0], cmdfn[1])
	switch cmdfn[0] {
	case "GET":
		//打开文件
		//读取内容
		//发送内容
		//关闭连接和文件
		f, err := os.Open(*root + cmdfn[1])
		if err != nil && err != io.EOF {
			log.Print(err)
			return
		}
		defer f.Close()
		// ct := bufio.NewReader(f)
		// for {
		// 	line, err := ct.ReadString('\n')
		// 	if err == io.EOF {
		// 		break
		// 	}
		// 	conn.Write([]byte(line))
		// }
		io.Copy(conn, f)
	case "STORE":
		// 从r读取文件内容直到err为io.EOF
		// 创建文件
		// 向文件写入内容
		// 往conn写入ok
		// 关闭连接和文件
		// os.MkdirAll(filepath.Dir(cmdfn[1]), 0755)
		f, _ := os.Create(*root + filepath.Base(cmdfn[1]))
		if err != nil {
			log.Print(err)
			return
		}
		// for {
		// 	line, err := r.ReadString('\n')
		// 	if err == io.EOF {
		// 		break
		// 	}
		// 	f.WriteString(line)
		// }
		io.Copy(f, r)
		f.Close()
		conn.Write([]byte("OK"))
	case "LS":
		f, err := os.Open(*root + cmdfn[1])
		if err != nil {
			log.Print(err)
		}
		defer f.Close()
		infos, _ := f.Readdir(-1)
		for _, info := range infos {
			conn.Write([]byte(fmt.Sprintf("%v %v \n", info.Name(), info.Size())))
		}
	default:
		fmt.Println("default")
		return
	}

}

func main() {
	addr := ":7777"                          //addr为IP:端口，一般忽略IP
	listener, err := net.Listen("tcp", addr) //监听端口
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close() //程序退出，关闭监听
	for {                  //循环接收请求
		conn, err := listener.Accept() //接收新的连接，阻塞连接
		if err != nil {
			log.Fatal(err)
		}
		go handConn(conn) //启动处理连接协程，哪里阻塞go哪里

	}
}
