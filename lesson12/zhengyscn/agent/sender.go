package main

import (
	"encoding/json"
	"fmt"
	"github.com/467754239/godoc/lesson/lesson12/monitor/common"
	"log"
	"net"
	"time"
)

type Sender struct {
	addr string
	ch   chan *common.Metric
}

func NewSender(addr string) *Sender {
	return &Sender{
		addr: addr,
		ch:   make(chan *common.Metric, 1),
	}
}

func (s *Sender) connect() net.Conn {
	n := 100 * time.Microsecond
	for {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil {
			log.Print(err)
			time.Sleep(n)
			n = n * 2
			if n > time.Second*30 {
				n = time.Second * 30
			}
			continue
		}
		return conn
	}
}

func (s *Sender) Start() {
	// 建立连接
	// 循环从s.ch里面读取metric
	// 序列化metric
	// 发送数据
	log.Print("before connect")
	conn := s.connect()
	log.Print("connect ok")
	log.Print(conn.LocalAddr().String())
	for metric := range s.ch {
		buf, _ := json.Marshal(metric)
		log.Print(metric)
		_, err := fmt.Fprintf(conn, "%s\n", string(buf))
		if err != nil {
			conn.Close()
			conn = s.connect()
			log.Print(conn.LocalAddr().String())
		}
	}
}

func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}
