package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gyf210/monitor/common"
	"log"
	"net"
	"time"
)

type Sender struct {
	ch   chan *common.Metric
	addr string
}

func NewSender(adder string) *Sender {
	return &Sender{
		ch:   make(chan *common.Metric, 1000),
		addr: adder,
	}
}

// 建立连接保持重连机制
func (s *Sender) connect() net.Conn {
	sleepTime := time.Duration(gconfig.Sender.MaxSleepTime) * time.Second
	n := 100 * time.Microsecond
	for {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil {
			log.Print(err)
			time.Sleep(n)
			n = n * 2
			if n > sleepTime {
				n = sleepTime
			}
			continue
		}
		return conn
	}
}

// 实现定时和定量发送数据保持相关的吞吐量
func (s *Sender) Start() error {
	conn := s.connect()
	w := bufio.NewWriter(conn)
	ticker := time.NewTicker(time.Duration(gconfig.Sender.FlushInterval) * time.Second)
	for {
		select {
		case metric := <-s.ch:
			buf, err := json.Marshal(metric)
			if err != nil {
				log.Print(err)
				continue
			}
			_, err = fmt.Fprintf(w, "%s\n", buf)
			if err != nil {
				conn.Close()
				conn = s.connect()
				log.Println(conn.LocalAddr())
				w = bufio.NewWriter(conn)
			}
		case <-ticker.C:
			err := w.Flush()
			if err != nil {
				conn.Close()
				conn = s.connect()
				log.Println(conn.LocalAddr())
				w = bufio.NewWriter(conn)
			}
		}
	}
	return nil
}

func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}

func (s *Sender) Close() {
	close(s.ch)
}
