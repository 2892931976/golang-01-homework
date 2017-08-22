package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/51reboot/golang-01-homework/lesson11/luojiyin/common"
	"github.com/51reboot/golang-01-homework/lesson12/luojiyin/monitor"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

var (
	addr = flag.String("trans", "59.110.12.72:6000", "transfer server")
)

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

func CpuMetric() []*commom.Metric {
	var ret []*commom.Metric
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
		metric = NewMetric("cpu.load15", cpuload.Load15)
		ret = append(ret, metric)
	}

	return ret
}

func main() {
	flage.Parse()
	sender := monitor.NewSender(*addr)
	ch := sender.Channel()

	scheder := monitor.NewSched(ch)
	go scheder.AddMetric(CpuMetric(), time.Second*2)
}
