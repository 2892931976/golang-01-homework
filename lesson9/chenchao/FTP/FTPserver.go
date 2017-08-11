package main

import (
	"bufio"
	"bytes"
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
	base_dir = flag.String("home", "base", "home dir") // 定义根目录
	//action = flag.String("a", "list", "Operation action")	// 具体操作
	//file_name = flag.String("n", "null", "file name")		// 操作文件名
	actionmap = map[string]func(filename string, logger *log.Logger, conn net.Conn){
		"list":   list,
		"get":    downfile,
		"upload": upload,
	}
)

// 获取目录下的所有文件名
func list(filename string, logger *log.Logger, conn net.Conn) {

	var files []string
	dir_list, err := ioutil.ReadDir(*base_dir)
	if err != nil {
		logger.Println("list file err" + err.Error())
		conn.Write([]byte("list file err" + err.Error()))
		return
	}
	for _, name := range dir_list {
		files = append(files, name.Name())
	}
	stringByte := strings.Join(files, "\n")

	ret := []byte(stringByte)
	conn.Write(ret)
	logger.Println("list all file")
	return

}

func downfile(filename string, logger *log.Logger, conn net.Conn) {

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// 文件不存在 返回错误
		logger.Println(filename + "not found.")
		conn.Write([]byte(filename + "not found"))
		return
	}
	f, err := os.Open(filename)
	if err != nil {
		logger.Println("download file err:" + err.Error())
		conn.Write([]byte("download file err:" + err.Error()))
		return
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f) // 返回的本身就是切片
	if err != nil {
		logger.Println([]byte("download file err:" + err.Error()))
		return
	}
	buf_s := []byte("")
	ok := []byte("OK")
	buf1 := [][]byte{ok, buf}

	conn.Write(bytes.Join(buf1, buf_s))
	return

}

func upload(filename string, logger *log.Logger, conn net.Conn) {
	fmt.Println("upload file is ", filename)

	conn.Write([]byte("OK")) // 返回获取成功的消息

	f, err := os.Create(filename)
	if err != nil {
		logger.Println("upload file err: ", err)
		return
	}
	defer f.Close()
	io.Copy(f, conn)
	logger.Println("upload file: ", filename)
	return

}

func handleconn(conn net.Conn) {

	// 分析参数
	// 打开文件
	// 发送文件
	// 关闭连接和文件
	// 定义客户端发来的命令行为: -a 操作类型  -n 操作的文件名。 共2部分
	defer conn.Close()

	logger, logfile := init_file() // 先初始化
	defer logfile.Close()

	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n') // 读取一行客户端发来的消息
	if err != nil {
		conn.Write([]byte("cmd err: " + err.Error()))
		log.Print(err)
		return
	}
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	cmd := fields[0]      // 操作
	filename := fields[1] // 文件名
	// 输入的参数已经获取了 现在需要根据参数 写对应的方法 get upload list 等等
	file_path := filepath.Join(*base_dir, filename) // 文件名
	filefunc := actionmap[cmd]
	if filefunc == nil {
		conn.Write([]byte("cmd argument err: not found arg " + cmd))
		return
	}
	filefunc(file_path, logger, conn) // 执行函数

}

// 创建或者打开日志文件 和 操作的日志对象
func create_log() (*log.Logger, *os.File) {
	var logfile *os.File
	var err error
	_, err = os.Stat("test.log")
	if os.IsNotExist(err) {
		fmt.Println("create log file ")
		// 日志文件不存在 直接创建
		logfile, err = os.Create("test.log")
		if err != nil {
			log.Fatal("ftp server init error: log file>", err)
		}

	} else {
		//logfile, err = os.Open("test.log")
		fmt.Println("open the log file ")
		logfile, err = os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("ftp server init error: log file>", err)
		}
	}

	logger := log.New(logfile, "", log.Ldate|log.Ltime)
	return logger, logfile
}

// 初始化一些东西  根目录  日志文件
func init_file() (*log.Logger, *os.File) {
	os.MkdirAll(*base_dir, 0755)    // 程序开始，先创建根目录，如果存在也不会有影响
	logger, logfile := create_log() // 创建日志对象与文件
	return logger, logfile
}

func main() {

	addr := ":10086"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() // 一开始想到了go这里 但是意义不大
		fmt.Println(conn.RemoteAddr(), "connect")
		if err != nil {
			log.Fatal(err)
		}
		go handleconn(conn)
	}
}
