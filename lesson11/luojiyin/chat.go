package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var gRoom = NewRoom()

const PASSWORD = "123456"

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
	log.Printf("[%s] logined", user)
}

func (r *Room) Leave(user string) {
	_, ok := r.users[user]
	if !ok {
		log.Printf("user: %s not logined in\n", user)
	} else {
		log.Printf("[%s] leaved ", user)
		r.users[user].Close()
		delete(r.users, user)
	}

}

func (r *Room) Broadcase(user, msg string) {
	log.Print("[broadcase] user num:", len(r.users))
	for name, conn := range r.users {
		if name == user {
			continue
		} else {
			log.Printf("[broadcase] user: %s \tmsg:%s", name, msg)
			msgInfo := fmt.Sprintf("%s: %s", user, msg)
			conn.Write([]byte(msgInfo))
		}
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

func handleConn(ch chan net.Conn) {
	for {
		conn := <-ch
		var usr, pwd, line string
		var rd *bufio.Reader
		retry := 1

		for ; retry <= 3; retry++ {
			conn.Write([]byte("\tplease input you name and pwd\n"))
			rd = bufio.NewReader(conn)
			line, _ = rd.ReadString('\n')
			fileds := strings.Fields(strings.TrimSpace(line))
			if len(fileds) != 2 {
				conn.Write([]byte("\tbad login info\n"))
				continue
			}

			usr, pwd = fileds[0], fileds[1]
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

func worker(user string, rd *bufio.Reader) {
	for {
		log.Printf("[worker] [%s] start to check", user)
		_, ok := gRoom.users[user]
		if !ok {
			break
		} else {
			log.Printf("[worker] [%s] waiting msg...", user)
			msg, err := rd.ReadString('\n')
			if err != nil {
				log.Printf("[worker] [%s] read with error:%s", user, err.Error())
				break
			}
			gRoom.Broadcase(user, msg)
			//log.Printf("%s say %s", user, msg)

		}
		//gRoom.Leave(user)
	}
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
