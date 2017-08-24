package main

import (
	"encoding/json"
	"fmt"
	"net"

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
		ch:   make(chan *common.Metric),
	}
}

func (s *Sender) Start() {
	// 建立连接
	// 循环从s.ch 里面读取metric
	// 序列化metric
	// 发送数据
	conn, err := net.Dial("tcp", s.addr)
	if err != nil {
		panic(err)
	}
	for metric := range s.ch {
		buf, _ := json.Marshal(metric)
		fmt.Fprintf(conn, "%s\n", buf)
	}
}

func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}
