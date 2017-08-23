package monitor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/luojiyin/common"
	"github.com/shirou/gopsutil/cpu"
	//	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
)

type Sender struct {
	addr string
	ch   chan *common.Metric
}

func NewMetric(metric string, value float64) *common.Metric {
	hostname, err := os.Hostname()
	if err != nil {
		log.Print(err)
	}
	return &common.Metric{
		Metric:    metric,
		Endpoint:  hostname,
		Value:     value,
		Tag:       []string{runtime.GOOS},
		Timestamp: time.Now().Unix(),
	}
}

func CpuMetric() []*common.Metric {
	var ret []*common.Metric
	cpus, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Print(err)
	}
	metric := NewMetric("cpu.usage", cpus[0])
	ret = append(ret, metric)

	cpuload, err := load.Avg()
	if err != nil {
		log.Print(err)
	} else {
		metric = NewMetric("cpu.load1", cpuload.Load1)
		ret = append(ret, metric)
		metric = NewMetric("cpu.load5", cpuload.Load5)
		ret = append(ret, metric)
		metric = NewMetric("cpuload15", cpuload.Load15)
		ret = append(ret, metric)
	}
	return ret
}

func NewSender(addr string) *Sender {
	return &Sender{
		addr: addr,
		ch:   make(chan *common.Metric, 1000), //提供1000的存储空间， 可以保留最近1000个数据
	}
}

func (s *Sender) connect() net.Conn {
	n := 100 * time.Microsecond
	for {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil { // 发生连接错误，不断尝试连接，间隔不断增长，最后30s一次
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
	conn := s.connect()
	w := bufio.NewWriter(conn)
	log.Print(conn.LocalAddr())
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case metric := <-s.ch:
			buf, _ := json.Marshal(metric)
			log.Print(string(buf))
			_, err := fmt.Fprintf(w, "%s\n", buf)
			if err != nil {

				conn.Close()
				conn = s.connect()
				w = bufio.NewWriter(conn)
				log.Print(conn.LocalAddr())

			}
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

//func main() {
//	flag.Parse()
//	//addr := "59.110.12.72:6000"
//	sender := NewSender(*addr)
//	go sender.Start()
//	ch := sender.Channel()
//
//	diskStat, err := disk.Usage("/dev/mapper/centos-root")
//	if err != nil {
//		log.Print(err)
//	}
//	log.Print(diskStat)
//	ticker := time.NewTicker(time.Second)
//	for range ticker.C {
//		//hostname, err := os.Hostname()
//		//if err != nil {
//		//	log.Print(err)
//		//}
//		//cpus, err := cpu.Percent(time.Second, false)
//		//if err != nil {
//		//	log.Print(err)
//		//}
//		//metric := &common.Metric{
//		//	Metric:    "cpu.usage",
//		//	Endpoint:  hostname,
//		//	Value:     cpus[0],
//		//	Tag:       []string{runtime.GOOS},
//		//	Timestamp: time.Now().Unix(),
//		//}
//		metric := getcpuUsage()
//
//		ch <- metric
//		buf, _ := json.Marshal(metric)
//		log.Println(string(buf))
//	}
//}

//func getcpuUsage() *common.Metric {
//	hostname, err := os.Hostname()
//	if err != nil {
//		log.Print(err)
//	}
//	cpus, err := cpu.Percent(time.Second, false)
//	if err != nil {
//		log.Print(err)
//	}
//	metric := &common.Metric{
//		Metric:    "cpu.usage",
//		Endpoint:  hostname,
//		Value:     cpus[0],
//		Tag:       []string{runtime.GOOS},
//		Timestamp: time.Now().Unix(),
//	}
//	return metric
//}
