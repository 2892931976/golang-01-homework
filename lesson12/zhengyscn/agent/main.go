package main

import (
	"os"
	"runtime"
	"time"

	"flag"

	"github.com/467754239/godoc/lesson/lesson12/monitor/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"log"
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

		metric = NewMetric("cpu.load5", cpuload.Load15)
		ret = append(ret, metric)
	}

	return ret
}

func DiskMetric() []*common.Metric {
	var ret []*common.Metric

	usage, err := disk.Usage("/")
	if err != nil {
		panic(err)
	}
	metric := NewMetric("disk.usesd_percent", usage.UsedPercent)
	ret = append(ret, metric)

	metric = NewMetric("disk.used", float64(usage.Used))
	ret = append(ret, metric)

	metric = NewMetric("disk.free", float64(usage.Free))
	ret = append(ret, metric)

	return ret
}

func MemMetric() []*common.Metric {
	var ret []*common.Metric

	vmem, err := mem.VirtualMemory()
	if err != nil {
		log.Print(err)
	}

	metric := NewMetric("mem.usesd_percent", vmem.UsedPercent)
	ret = append(ret, metric)

	metric = NewMetric("mem.used", float64(vmem.Used))
	ret = append(ret, metric)

	metric = NewMetric("mem.free", float64(vmem.Free))
	ret = append(ret, metric)

	metric = NewMetric("mem.available", float64(vmem.Available))
	ret = append(ret, metric)

	return ret
}

func main() {
	flag.Parse()
	// 初始化构造函数
	sender := NewSender(*transAddr)
	// 返回构造函数的ch
	ch := sender.Channel()

	sched := NewSched(ch)

	// cpu, time.Second * 1
	// memory, time.Second * 3
	// disk, time.Minute
	go sched.AddMetric(CpuMetric, time.Second*1)
	go sched.AddMetric(DiskMetric, time.Minute)
	go sched.AddMetric(MemMetric, time.Second*3)

	sender.Start()

}
