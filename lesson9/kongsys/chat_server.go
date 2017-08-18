package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var globalRoom *Room

type Room struct {
	users map[string]net.Conn
}

func NewRoom() *Room {
	return &Room{
		users: make(map[string]net.Conn),
	}
}
func (r *Room) Broadcast(user string, msg string) {
	for u, v := range r.users {
		if u == user {
			continue
		}
		fmt.Fprintf(v, "%s:%s", user, msg)
	}
}

func (r *Room) Leave(user string) {
	conn, ok := r.users[user]
	if ok {
		conn.Close()
	}
	delete(r.users, user)
}

func (r *Room) Join(user string, conn net.Conn) {
	_, ok := r.users[user]
	if ok {
		r.Leave(user)
	}
	r.users[user] = conn
}
func handleConn(conn net.Conn) {
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len(fields) != 2 {
		conn.Write([]byte("bad input\n"))
		return
	}
	user := fields[0]
	password := fields[1]
	if password != "123456" {
		conn.Write([]byte("error password!"))
		return
	}
	globalRoom.Join(user, conn)
	fmt.Fprintf(conn, "Hi, %s\n", user)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			globalRoom.Leave(user)
			break
		}
		globalRoom.Broadcast(user, line)
	}

}

func main() {
	globalRoom = NewRoom()
	addr := ":8021"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
