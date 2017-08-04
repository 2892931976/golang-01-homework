/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	cmd  string
	line string
)

func main() {
	addr := "172.16.3.210:8080"        //定义主机名
	conn, err := net.Dial("tcp", addr) //拨号操作，用于连接服务端，需要指定协议。
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 10240)       //定义一个切片的长度是10240。
	n, err := conn.Read(buf)         //接收到的内容大小为我们提前定义好的大小。
	if err != nil && err != io.EOF { //io.EOF在网络编程中表示对端把链接关闭了。
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n])) //将接受的内容都读取出来。

	f := bufio.NewReader(os.Stdin)
	for {
		//fmt.Print("请输入>>>:")
		line, err = f.ReadString('\n') //定义一行的内容，结束标识符是换行符“\n”
		fmt.Sscan(line, &cmd)
		if len(line) == 1 {
			continue
		}
		//fmt.Print(line)
		go sender(conn, line)
	}
	conn.Close() //断开TCP链接。
}

func sender(conn net.Conn, line string) {
	n, err := conn.Write([]byte(line)) //向服务端发送数据。用n接受返回的数据大小，用err接受错误信息。
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 10) //定义一个切片的长度是1024。

	for {
		n, err = conn.Read(buf) //接收到的内容大小。
		if err == io.EOF {
			conn.Close()
		}
		fmt.Print(string(buf[:n]))
	}
	return
}
