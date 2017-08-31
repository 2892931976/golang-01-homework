package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/51reboot/golang-01-homework/lesson13/jungle85gopy/common"
)

// Sender for agent
type Sender struct {
	addr          string
	maxRetry      int
	flushInterval int
	ch            chan *common.Metric
}

// NewSender 构造函数
func NewSender(scfg SenderConfig) *Sender {
	sender := &Sender{
		addr:          scfg.TransAddr,
		maxRetry:      scfg.MaxSleepTime,
		flushInterval: scfg.FlushInterval,
		ch:            make(chan *common.Metric, scfg.BufferLen),
	}
	return sender
}

// Channel 返回ch
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
			if baseGap > time.Second*time.Duration(s.maxRetry) {
				baseGap = time.Second * time.Duration(s.maxRetry)
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

	ticker := time.NewTicker(time.Second * time.Duration(s.flushInterval))
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
