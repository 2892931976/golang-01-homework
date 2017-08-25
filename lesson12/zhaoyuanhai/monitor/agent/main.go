package main

import (
	"flag"
	"lesson12/monitor/common"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

var (
	transaddr = flag.String("trans", "59.110.12.72:6000", "transfer address")
)

func NewMetric(metric string, value float64) *common.Metric {
	hostname, _ := os.Hostname()
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
	// hostname, _ := os.Hostname()
	cpus, err := cpu.Percent(time.Second, false)
	if err != nil {
		panic(err)
	}
	metric := NewMetric("cpu.usage", cpus[0])
	ret = append(ret, metric)
	cpuload, err := load.Avg()
	if err == nil {
		metric = NewMetric("cpu.load1", cpuload.Load1)
		ret = append(ret, metric)
		metric = NewMetric("cpu.load5", cpuload.Load5)
		ret = append(ret, metric)
	}
	return ret
}

func MemMetric() []*common.Metric {
	var ret []*common.Metric
	memstat, err := mem.VirtualMemory()
	if err != nil {
		log.Println(err)
	}
	metric := NewMetric("mem.usage", memstat.UsedPercent)
	ret = append(ret, metric)
	return ret
}

func main() {

	flag.Parse()
	sender := NewSender(*transaddr)
	ch := sender.Channel()
	// ticker := time.NewTicker(time.Second * 5)
	// for range ticker.C {

	// 	cpus, err := cpu.Percent(time.Second, false)

	// 	fmt.Println(metric)
	// 	// 生产者---》消费者模式之生产者
	// 	ch <- metric
	// }
	sched := NewSched(ch)
	// log.Print("main: CpuMetric", CpuMetric)
	go sched.AddMetric(CpuMetric, time.Second*5)
	go sched.AddMetric(MemMetric, time.Second*5)
	sender.Start()
}
