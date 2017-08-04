package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
	"io"
	"time"
)
var globalRoom *Room = NewRoom()

type Room struct {
	users map[string]net.Conn
}

func NewRoom() *Room {
	return &Room{
		users: make(map[string]net.Conn),
	}
}

func (r *Room) Join(user string, conn net.Conn)  {
	if r.users[user] != nil{
		r.Leave(user)
	}
	r.users[user] = conn
	fmt.Printf("%s 登录成功。", user)
}

func (r *Room) Leave(user string)  {
	r.users[user].Close()
	delete(r.users, user)
	fmt.Printf("%s 离开", user)
}

func (r *Room) Broadcast(user string, msg string) {
	// 哪个用户 发送了消息
	for _, v := range r.users{
		v.Write([]byte(user+": "+msg))
	}
}


func Work(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')
	if err != nil{
		fmt.Println(err)
		return
	}
	cmd := strings.Fields(line)
	fmt.Println(strings.Fields(line))
	if cmd[0] == "login"{
		if cmd[2] != "123456"{
			return
		}
		globalRoom.Join(cmd[1], conn)
		conn.Write([]byte("登录成功"))
		for {		// 开始接受发送的消息
			msg, err :=r.ReadString('\n')
			if err != nil{
				return
			}else if err == io.EOF {
				return
			}
			fmt.Println(msg)
			globalRoom.Broadcast(cmd[1], msg)
			time.Sleep(time.Second * 15)
		}
		// leave 用户
		globalRoom.Leave(cmd[1])
		conn.Write([]byte("退出成功"))
	}

}

func main() {
	addr := "0.0.0.0:8033"
	listener, err := net.Listen("tcp", addr)
	if err != nil{
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go Work(conn)

	}
}
