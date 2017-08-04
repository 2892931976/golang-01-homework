package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
	"io"
	//"time"
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
	conn.Write([]byte(user+ ":加入聊天室!\n"))
}

func (r *Room) Leave(user string)  {
	r.users[user].Close()
	delete(r.users, user)
	fmt.Printf("%s 离开", user)
}

func (r *Room) Broadcast(user string, msg string) {
	// 哪个用户 发送了消息
	fmt.Println("广播：",user)
	//user = strings.TrimSpace(user)
	fmt.Println("信息：",msg)
	msg = strings.TrimSpace(msg)
	fmt.Println(len(msg))
	for _, v := range r.users{
		v.Write([]byte(user+":"+msg))
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
	if len(cmd) < 2 {
		conn.Write([]byte("您输入的字符串用户名活密码无效，程序强制退出！\n"))
		return
	}
	if cmd[1] == "123"{
		globalRoom.Join(cmd[0], conn)
		conn.Write([]byte("登录成功\n"))
		for {		// 开始接受发送的消息
			msg, err :=r.ReadString('\n')
			if err != nil{
				return
			}else if err == io.EOF {
				return
			}
			fmt.Println(msg)
			globalRoom.Broadcast(cmd[0], msg)
			fmt.Println()
			//time.Sleep(time.Second * 15)
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

