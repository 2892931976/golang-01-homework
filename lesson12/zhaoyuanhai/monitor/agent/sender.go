package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"lesson12/monitor/common"
	"log"
	"net"
	"time"
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
		log.Print(conn.LocalAddr())
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
	//循环从s.ch里面读取Metric
	//序列化metric
	//发送数据
	// conn, err := net.Dial("tcp", s.addr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer conn.Close()
	conn := s.connect()
	w := bufio.NewWriter(conn)
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case metric := <-s.ch:
			buf, _ := json.Marshal(metric)
			fmt.Println("conn: ", conn)
			_, err := fmt.Fprintf(conn, "%s\n", buf)
			if err != nil {
				conn.Close()
				conn = s.connect()
			}
		case <-ticker.C:
			err := w.Flush()
			if err != nil {
				conn.Close()
				conn = s.connect()
			}
		}
	}
	// for metric := range s.ch {
	// 	buf, _ := json.Marshal(metric)
	// 	_, err := fmt.Fprintf(conn, "%s\n", buf)
	// 	if err != nil {
	// 		conn.Close()
	// 		conn = s.connect()
	// 	}

	// }
	// for {
	// 	//生产者---》消费者模式之消费者
	// 	metric := <-s.ch
	// 	buf, _ := json.Marshal(metric)
	// 	//向conn里发送数据
	// 	_, err := fmt.Fprintf(conn, "%s\n", buf)
	// 	if err != nil {
	// 		conn.Close()
	// 		conn = s.connect()
	// 	}
	// }
	// 方法二
	// for metric := range s.ch {
	// 	buf, _ := json.Marshal(metric)
	// 	//向conn里发送数据
	// 	fmt.Fprintf(conn, "%s\n", buf)
	// }
}

func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}
