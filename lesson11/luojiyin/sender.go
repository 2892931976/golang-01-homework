package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/luojiyin/monitor/common"
	"github.com/shirou/gopsutil/cpu"
)

type Sender struct {
	addr string
	ch   chan *common.Metric
}

func NewSender(addr string) *Sender {
	return &Sender{
		addr: addr,
		ch:   make(chan *common.Metric),
	}
}
func (s *Sender) Start() {
	conn, err := net.Dial("tcp", s.addr)
	if err != nil {
		panic(err)
	}
	for {
		metric := <-s.ch
		buf, err := json.Marshal(metric)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(conn, "%s\n", buf)

	}

}

func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}

func main() {
	hostname, _ := os.Hostname()
	cpus, err := cpu.Percent(time.Second, false)
	if err != nil {
		panic(err)
	}
	metric := &common.Metric{
		Metric:    "cpu.usage",
		Endpoint:  hostname,
		Value:     cpus[0],
		Tag:       []string{runtime.GOOS},
		Timestamp: time.Now().Unix(),
	}
	buf, _ := json.Marshal(metric)
	fmt.Println(string(buf))

	addr := "59.110.12.72:6000"
	sender := NewSender(addr)

	go sender.Start()
	ch := sender.Channel()

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		hostname, _ := os.Hostname()
		cpus, err := cpu.Percent(time.Second, false)
		if err != nil {
			panic(err)
		}
		metric := &common.Metric{
			Metric:    "cpu.usage",
			Endpoint:  hostname,
			Value:     cpus[0],
			Tag:       []string{runtime.GOOS},
			Timestamp: time.Now().Unix(),
		}

		ch <- metric
	}
}
