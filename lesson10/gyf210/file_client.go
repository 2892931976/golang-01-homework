package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	for {
		addr := "127.0.0.1:9000"
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatalln(err)
		}
		s := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		line, _ := s.ReadString('\n')
		cmd := strings.Fields(line)
		if len(cmd) == 0 {
			continue
		}
		switch cmd[0] {
		case "GET":
			if len(cmd) != 2 {
				fmt.Println(">> [GET a.txt]")
				break
			}
			conn.Write([]byte(line))
			f, err := os.Create(cmd[1])
			if err != nil {
				log.Fatalln(err)
			}
			defer f.Close()
			io.Copy(f, conn)
		case "LS":
			if len(cmd) > 2 {
				fmt.Println(">> [LS | LS /a/b/c]")
				break
			}
			conn.Write([]byte(line))
			io.Copy(os.Stdout, conn)
		case "STORE":
			if len(cmd) != 3 {
				fmt.Println(">> [STORE a.txt b.txt]")
				break
			}
			t := fmt.Sprintf("%v %v\n", cmd[0], cmd[1])
			f, err := os.Open(cmd[2])
			if err != nil {
				fmt.Println(err)
				break
			}
			defer f.Close()
			conn.Write([]byte(t))
			io.Copy(conn, f)
		case "QUIT":
			conn.Close()
			return
		default:
			fmt.Println(">> [GET | LS | STORE | QUIT]")
		}
		conn.Close()
	}
}
