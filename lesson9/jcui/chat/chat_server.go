package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

/*
监听端口
接受新的连接
启动协程
发送接收数据
断开连接
*/
type Room struct {
	users map[string]net.Conn
}

func NewRoom() *Room {
	return &Room{
		users: make(map[string]net.Conn),
	}
}

func (r *Room) Join(user string, conn net.Conn) {
	_, ok := r.users[user]
	if ok {
		r.Leave(user)
	}
	r.users[user] = conn
}

func (r *Room) Leave(user string) {
	//关闭连接
	//从users里面删除
	//r.users[user].Close()
	//delete(r.users, user)
	//
	conn, ok := r.users[user]
	if !ok {
		return
	}
	conn.Close()
	delete(r.users, user)
}

func (r *Room) Broadcast(user string, msg string) {
	//r.users[user].Read([]byte(msg))

	for name, conn := range r.users {
		if name == user {
			continue
		}
		tosend := fmt.Sprintf("%s:%s\n", user, msg)
		conn.Write([]byte(tosend))
	}
}

/*
client -> binggan 123456
client -> hello golang
client -> close

接收新的连接
验证用户的账号和密码
等待用户输入内容
广播所有在线的用户广播用户的输入


*/
func handleConn(conn net.Conn) { //主机conn 这里的类型net.Conn
	defer conn.Close()
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len(fields) != 2 {
		conn.Write([]byte("bad input"))
		return
	}
	user := fields[0]
	password := fields[1]
	if password != "123456" {
		return
	}
	fmt.Println(user, password)
	// Join用户
	NewRoom().Join(user, conn)
	//广播一下用户登录信息
	NewRoom().Broadcast("SYSTEM:", fmt.Sprintf("%s join Room", user))
	//用户聊天
	for {

		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		//broadcast
		NewRoom().Broadcast(user, line)

	}
	//
	NewRoom().Leave(user)
	NewRoom().Broadcast("SYSTEM", fmt.Sprintf("%s leave Room", user))

}

func main() {
	addr := ":7777" //监听任意IP的7777端口
	//创建监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	//
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)

	}

}
