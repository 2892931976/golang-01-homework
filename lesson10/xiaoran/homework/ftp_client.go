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

func main() {
	addr := "127.0.0.1:8021"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("连接ftp成功，请输入命令")
	defer conn.Close()

	f := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, _ := f.ReadString('\n')
		// 去除两端的空格和换行
		line = strings.TrimSpace(line)
		// 按空格分割字符串得到字符串列表
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		// 获取命令和参数列表
		cmd := args[0]
		name := args[1]
		switch cmd {
		case "GET":
			ftp_cmd := line + "\n"
			_, err := conn.Write([]byte(ftp_cmd))
			if err != nil {
				log.Fatal(err)
			}
			if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
				if os.IsPermission(err) {
					fmt.Println("你不够权限")
				}
			}

			f, err := os.Create(name)
			if err != nil {
				log.Print(err)
				return
			}
			defer f.Close()
			io.Copy(f, conn)

		case "STORE":
			ftp_cmd := line + "\n"
			_, err := conn.Write([]byte(ftp_cmd))
			if err != nil {
				log.Fatal(err)
			}
			f, err := os.Open(name)
			if err != nil {
				log.Print(err)
				return
			}
			defer f.Close()
			io.Copy(conn, f)

		case "LS":
			ftp_cmd := line + "\n"
			_, err := conn.Write([]byte(ftp_cmd))
			if err != nil {
				log.Fatal(err)
			}
			io.Copy(os.Stdout, conn)

		default:
			fmt.Println("请输入:GET|STORE|LS")
		}
	}
}
