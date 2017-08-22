package main

import (
	"flag"
	"github.com/gyf210/monitor/agent"
	"github.com/gyf210/monitor/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	target   = flag.String("t", "", "transfer address")
	hostname = GetHostName()
)

func GetHostName() (hostname string) {
	hostname, _ = os.Hostname()
	return
}

func NewMetric(metric string, value float64) *common.Metric {
	return &common.Metric{
		Metric:    metric,
		Endpoint:  hostname,
		Value:     value,
		Tag:       []string{runtime.GOOS},
		Timestamp: time.Now().Unix(),
	}
}

func MemMetric() []*common.Metric {
	var ret []*common.Metric
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Println(err)
		return nil
	}
	metric := NewMetric("mem.usage", m.UsedPercent)
	ret = append(ret, metric)
	return ret
}

func CpuMetric() []*common.Metric {
	var ret []*common.Metric
	c, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Println(err)
		return nil
	}
	metric := NewMetric("cpu.usage", c[0])
	ret = append(ret, metric)

	cpuload, err := load.Avg()
	if err != nil {
		log.Println(err)
		return nil
	}
	metric = NewMetric("cpu.load1", cpuload.Load1)
	ret = append(ret, metric)
	metric = NewMetric("cpu.load5", cpuload.Load5)
	ret = append(ret, metric)
	return ret
}

func main() {
	flag.Parse()
	sender := agent.NewSender(*target)
	ch := sender.Channel()

	scheder := agent.NewSched(ch)
	go scheder.AddMetric(CpuMetric, time.Second*5)
	go scheder.AddMetric(MemMetric, time.Second*2)

	sender.Start()
}
