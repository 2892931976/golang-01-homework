package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"learning/mygo_local/lesson12/monitor/common"
	"log"
	"net"
	"time"
)

type Sender struct {
	addr string
	ch   chan *common.Metric
}

func NewSender(addr string) *Sender {
	//构造sende
	return &Sender{
		addr: addr,
		ch:   make(chan *common.Metric, 10000),
	}
}
func (s *Sender) connect() net.Conn {
	n := 100 * time.Millisecond
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
	//建立连接
	//循环从s.sh里面读取metric
	//序列化metric
	//发送数据
	conn, err := net.Dial("tcp", s.addr)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(conn)
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case metric := <-s.ch:
			buf, _ := json.Marshal(metric)
			_, err := fmt.Fprintln(w, "%s\n", buf)
			if err != nil {
				conn.Close()
				conn = s.connect()
				w := bufio.NewWriter(conn)
				log.Print(w)
			}
		case <-ticker.C:
			w.Flush()

		}

	}
	/*
		for metric := range s.ch {
			buf, _ := json.Marshal(metric)
			_, err := fmt.Fprintln(conn, "%s\n", buf)
			if err != nil {
				conn.Close()
				conn = s.connect()
			}
		}
	*/
}

func (s *Sender) Channel() chan *common.Metric {

	return s.ch
}
