package main

import (
	"monitor/common"
	"net"
	"log"
	"fmt"
	"encoding/json"
	"time"
	"bufio"
)

type Sender struct {
	addr string
	ch chan *common.Metric
}

// 初始化一个可以存放数据的chan
func NewSender(add string) *Sender {
	ch := make(chan *common.Metric, 1000)		// 加1000后具有缓存的功能
	return &Sender{addr:add, ch: ch}

}

func (s *Sender) reConnect() net.Conn {
	fmt.Println("create connect ....")
	n := 100 * time.Millisecond
	for {
		fmt.Println("re connect...")
		conn, err := net.Dial("tcp", s.addr)
		if err != nil{
			log.Println(err)
			time.Sleep(n)
			n = n * 2
			if n > 30 * time.Second{
				n = time.Second * 30
			}
			continue
		}
		return conn
	}
}

func (s *Sender) Start() {
	// 循环的从chan里读取数据 发送给客户端
	fmt.Println("Start...")
	server := s.reConnect()
	buf_server := bufio.NewWriter(server)
	ticker := time.NewTicker(time.Second * 3)

	for {
		select {
		case metric := <-s.ch:
			buf, _ := json.Marshal(metric)
			_, err := fmt.Fprintf(buf_server, "%s\n", buf)
			if err != nil{
				server.Close()
				server = s.reConnect()
				buf_server = bufio.NewWriter(server)

			}
		case  <- ticker.C:
			err := buf_server.Flush()
			if err != nil{
				server.Close()
				server = s.reConnect()
				buf_server  = bufio.NewWriter(server)

			}
		}
	}

}

func (s *Sender) Channel() chan *common.Metric{
	fmt.Println("Channel...")
	return s.ch
}
