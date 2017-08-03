package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	addr := "127.0.0.1:7777"
	//拨号,请求连接
	conn, err := net.Dial("tcp", addr)
	//如果没有错误,则表示连接成功
	if err != nil {
		log.Fatal(err)
	}
	//记得关闭
	defer conn.Close()
	//打开文件
	fd, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	//fmt.Println(conn.RemoteAddr().String()) //打印远端地址及端口
	//fmt.Println(conn.LocalAddr().String())  //打印本地地址及建立连接所随机的端口

	//发送数据
	n, err := conn.Write([]byte(os.Args[1] + " " + filepath.Base(os.Args[1]) + "\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n) //代表发送了18个字节

	/*
		如下这段内容参考同学作业,一直卡在这里不返回,收不到EOF 原因是io>copy不发送EOF
		这里使用断言，关闭发送，以便发送EOF
	*/
	io.Copy(conn, fd)
	v, _ := conn.(*net.TCPConn)
	v.CloseWrite()

	//读取返回结果
	//buf := make([]byte, 4096)        //创建一个缓冲区
	//n, err = conn.Read(buf)          //n代表读取的行数
	//if err != nil && err != io.EOF { //读取到结尾,EOF代表对方把连接关闭
	//	log.Fatal(err)
	//}
	//fmt.Println(n, string(buf[:n]))
	fmt.Println("Complate")

}
