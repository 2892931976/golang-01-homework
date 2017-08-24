package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/wujianjian/monitor/common"
)

type Sender struct {
	addr string
	ch   chan *common.Metric
}

func NewSender(addr string) *Sender {
	//构造sender
	return &Sender{
		addr: addr,
		ch:   make(chan *common.Metric, 10000),
	}
}

func (s *Sender) connect() net.Conn {
	n := time.Millisecond * 100
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
	// 循环从s.ch 里面读取metric
	// 序列化metric
	// 发送数据
	conn := s.connect()
	log.Print(conn.LocalAddr())
	w := bufio.NewWriter(conn)
	ticker := time.NewTicker(time.Second * 5)
	for {
		// 定时定量发送
		select {
		case metric := <-s.ch:
			buf, _ := json.Marshal(metric)
			_, err := fmt.Fprintf(w, "%s\n", buf)
			if err != nil {
				conn.Close()
				conn = s.connect()
				w = bufio.NewWriter(conn)
				log.Print(conn.LocalAddr())
			}
		case <-ticker.C:
			err := w.Flush()
			if err != nil {
				conn.Close()
				conn = s.connect()
				w = bufio.NewWriter(conn)
				log.Print(conn.LocalAddr())
			}
		}
	}
}

func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}
