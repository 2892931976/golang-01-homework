package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var gRoom = NewRoom()

// PASSWORD for all user password
const PASSWORD = "123456"

// Room for user map
type Room struct {
	users map[string]net.Conn
}

// NewRoom new a room
func NewRoom() *Room {
	return &Room{
		users: make(map[string]net.Conn),
	}
}

// Join for user join into room
func (r *Room) Join(user string, conn net.Conn) {
	_, ok := r.users[user]
	if ok {
		r.Leave(user)
	}
	r.users[user] = conn
	log.Printf("[%s] loggined", user)
}

// Leave for user leave from room
// close conn; delete user
func (r *Room) Leave(user string) {
	_, ok := r.users[user]
	if !ok {
		log.Printf("user: %s not logined in\n", user)
	}
	log.Printf("[%s] leaved", user)
	r.users[user].Close()
	delete(r.users, user)
}

// Broadcase broadcase msg to all user in a room
// 遍历所有的用户。发送消息
func (r *Room) Broadcase(user, msg string) {
	log.Print("[broadcase] user num:", len(r.users))
	for name, conn := range r.users {
		if name == user {
			continue
		}
		log.Printf("[broadcase] user: %s\tmsg:%s", name, msg)
		msgInfo := fmt.Sprintf("%s: %s", user, msg)
		conn.Write([]byte(msgInfo))
	}
}

func server(listener net.Listener) {
	connCh := make(chan net.Conn)
	go handleConn(connCh)
	go handleConn(connCh)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		connCh <- conn
	}
}

/*
client : nc localhost 8021
client -> bingan 123456
client -> hello golang
client -> close
*/
// 接收新的连接。验证用户的账户和密码；加入聊天室；
func handleConn(ch chan net.Conn) {
	for {
		conn := <-ch
		var usr, pwd, line string
		var rd *bufio.Reader
		retry := 1
		for ; retry <= 3; retry++ {
			conn.Write([]byte("\tpls input you name and pwd\n"))
			rd = bufio.NewReader(conn)
			line, _ = rd.ReadString('\n')
			fields := strings.Fields(strings.TrimSpace(line))
			if len(fields) != 2 {
				conn.Write([]byte("\tbad login info\n"))
				continue
			}
			usr, pwd = fields[0], fields[1]
			if pwd == PASSWORD {
				break
			}
			conn.Write([]byte("\tbad password\n"))
			log.Print("bad password")
		}
		if retry > 3 {
			conn.Close()
			continue
		}
		gRoom.Join(usr, conn)
		conn.Write([]byte("\tlogin success\n"))
		go worker(usr, rd)
	}
}

// 等待用户输入；向所有在线用户广播用户的输入
func worker(user string, rd *bufio.Reader) {
	for {
		log.Printf("[worker] [%s] start to check", user)
		_, ok := gRoom.users[user]
		if !ok {
			break
		}
		log.Printf("[worker] [%s] waiting msg...", user)
		msg, err := rd.ReadString('\n')
		if err != nil {
			log.Printf("[worker] [%s] read with error:%s", user, err.Error())
			break
		}
		gRoom.Broadcase(user, msg)
	}
	// leave user
	gRoom.Leave(user)
}

func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Print(err)
	}
	defer listener.Close()

	server(listener)
}
