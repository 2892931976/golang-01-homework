package main

import (
	"flag"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"learning/mygo_local/lesson12/monitor/common"
	"os"
	"runtime"
	"time"
)

var (
	transAddr = flag.String("trans", "59.110.12.72:6000", "transfef addr")
)

func NewMetric(metric string, value float64) *common.Metric {
	hostname, _ := os.Hostname()
	return &common.Metric{
		Metric:    metric,
		Endpoint:  hostname,
		Tag:       []string{runtime.GOOS},
		Value:     value,
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
	//

	return ret
}

func main() {
	flag.Parse()
	sender := NewSender(*transAddr)
	//go sender.Start()
	ch := sender.Channel()
	/*
		ticker := time.NewTicker(time.Second * 5)

		for range ticker.C {

			hostname, _ := os.Hostname()
			cpus, err := cpu.Percent(time.Second, false)
			if err != nil {
				panic(err)
			}
			metric := &common.Metric{ //引入common包中的结构体
				Metric:    "cpu.usage",
				Endpoint:  hostname,
				Tag:       []string{runtime.GOOS},
				Value:     cpus[0],
				Timestamp: time.Now().Unix(),
			}
			fmt.Println(metric)
			ch <- metric
		}
	*/

	sched := NewSched(ch)
	sched.AddMetric(CpuMetric, time.Second*5)
	sender.Start()

}
