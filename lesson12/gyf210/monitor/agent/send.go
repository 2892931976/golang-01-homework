package agent

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
	addr string
	ch   chan *common.Metric
}

func NewSender(addr string) *Sender {
	return &Sender{
		addr: addr,
		ch:   make(chan *common.Metric, 10000),
	}
}

func (s *Sender) connect() net.Conn {
	n := 100 * time.Microsecond
	for {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil {
			log.Println(err)
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
	conn := s.connect()
	w := bufio.NewWriter(conn)
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case m := <-s.ch:
			b, _ := json.Marshal(m)
			_, err := fmt.Fprintf(w, "%s\n", b)
			if err != nil {
				conn.Close()
				conn = s.connect()
				w = bufio.NewWriter(conn)
				log.Print(conn.LocalAddr())
			}
			log.Println(string(b))
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
