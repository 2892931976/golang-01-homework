package main

import (
	"flag"
	"os"
	"runtime"
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/wujianjian/monitor/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

var (
	transAddr = flag.String("trans", "59.110.12.72:6000", "transfer address")
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

func main() {
	flag.Parse()
	sender := NewSender(*transAddr)

	go sender.Start()
	ch := sender.Channel()

	sched := NewSched(ch)

	sched.AddMetric(CpuMetric, time.Second)
	// memory , time.Second * 3
	// disk , time.Minute
	sender.Start()
}
