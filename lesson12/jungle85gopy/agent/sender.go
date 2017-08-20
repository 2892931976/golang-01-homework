package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/jungle85gopy/common"
)

// Sender for agent
type Sender struct {
	addr string
	ch   chan *common.Metric
}

// NewSender new Sender
func NewSender(addr string) *Sender {
	sender := &Sender{
		addr: addr,
		ch:   make(chan *common.Metric, 1024),
	}
	return sender
}

// Channel get chan
func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}

// connect retry connect to transfer.
func (s *Sender) connect() net.Conn {
	baseGap := 500 * time.Millisecond
	for {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil {
			log.Print(err)
			time.Sleep(baseGap)
			baseGap *= 2
			if baseGap > time.Second*30 {
				baseGap = time.Second * 30
			}
			continue
		}
		debugInfo(fmt.Sprintf("local addr:%s\n", conn.LocalAddr()))
		return conn
	}
}

// reConnect retry connect while write to conn err
func (s *Sender) reConnect(conn net.Conn) *bufio.Writer {
	conn.Close()
	conn = s.connect()
	w := bufio.NewWriter(conn)
	return w
}

// Start 建立连接；
// 循环从ch中读取metric，序列化metric，发送数据
func (s *Sender) Start() {
	var conn net.Conn
	conn = s.connect()
	w := bufio.NewWriter(conn)

	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case metric := <-s.ch:
			buf, _ := json.Marshal(metric)
			debugInfo("~~ sender get metric")
			_, err := fmt.Fprintf(w, "%s\n", buf)
			if err != nil {
				log.Printf("Fprintf to remote err:%s", err.Error())
				w = s.reConnect(conn)
			}
		case <-ticker.C:
			debugInfo("-- Flush data to transfer from bufio of conn.")
			err := w.Flush()
			if err != nil {
				log.Printf("Flush to remote err:%s", err.Error())
				w = s.reConnect(conn)
			}
		}
	}
}
