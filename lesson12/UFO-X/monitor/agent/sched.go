package agent

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang-01-homework/lesson12/UFO-X/monitor/common"
)

type Sched struct {
	addr string
	ch   chan *common.Metric
}

func NewSched(addr string, ch chan *common.Metric) *Sched {
	return &Sched{
		addr: addr,
		ch:   ch,
	}
}
func (s Sched) Send() {
	tick := time.NewTicker(15 * time.Second)
	conn := s.connect()
	buf := bufio.NewWriter(conn)
	for {

		select {
		case metric := <-s.ch:
			js, _ := json.Marshal(metric)
			_, err := fmt.Fprintf(buf, "%s\n", js)
			if err != nil {
				log.Println(err)
				conn.Close()
				conn = s.connect()
				buf = bufio.NewWriter(conn)

			}

		case <-tick.C:
			fmt.Println("tick")
			err := buf.Flush()
			if err != nil {
				log.Println(err)
				conn.Close()
				conn = s.connect()
				buf = bufio.NewWriter(conn)
			}

		}

	}
}
func (s Sched) connect() net.Conn {

	tm := 100 * time.Millisecond
	for {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil {
			log.Println(err)
			time.Sleep(tm)
			tm = tm * 2
			if tm > 30*time.Second {
				tm = 30 * time.Second
			}
			continue
		}
		return conn
	}

}
