/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/


package main


import (
	"net"
	"log"
	"bufio"
	"os"
	"io"
	"fmt"
	"strings"
	"path/filepath"
)

func main() {
	addr := "172.16.3.211:8080"
	conn,err := net.Dial("tcp",addr)
	if err != nil{
		log.Fatal(err)
	}

	f:=bufio.NewReader(os.Stdin)
	for{
		fmt.Print("请输入:>>>")
		line,err := f.ReadString('\n')
		if err != nil {
			break
		}
		input := strings.Fields(line)
		if len(input) != 2 {
			fmt.Println("Usage: STORE|GET filename")
			continue
		}
		Sender(conn,line)
	}
}

func Sender_value(conn net.Conn ,msg string)  {
	f,err := os.Open(msg)
	if err != nil {
		panic("打开文件出错！")
	}
	defer f.Close()
	io.Copy(conn,f) //将消息发送给conn
	conn.Write([]byte("\n"))
}


func Sender(conn net.Conn,msg string)  {
	defer 	conn.Close()
	r := bufio.NewReader(conn)
	conn.Write([]byte(msg))
	cmd := strings.Fields(msg)[0]
	file_name := strings.Fields(msg)[1]
	switch cmd {
	case "GET", "get":
		fmt.Println("开始读取文件！")
		os.MkdirAll(filepath.Dir(file_name), 0755)
		f, err := os.Create(file_name)
		if err != nil {
			log.Print(err)
			return
		}
		io.Copy(f, r) //讲链接读取到的内容写入到刚刚创建的文件中。
		defer  f.Close()
		fmt.Println("文件读取完毕！")
	case "STORE","store":
		fmt.Println("开始发送文件！")
		Sender_value(conn,file_name)
		fmt.Println("文件发送完毕！")
	}
}

