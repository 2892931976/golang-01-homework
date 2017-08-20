package main

import (
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/jungle85gopy/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// SPtr2Metric slice ptr to metric
type SPtr2Metric []*common.Metric

// CPUMetric for CPU
func CPUMetric() (ret SPtr2Metric) {
	cpus, err := cpu.Percent(time.Second, false)
	if err != nil {
		panic(err)
	}
	metric := common.NewMetric("cpu.usage", cpus[0])
	ret = append(ret, metric)

	cpuload, err := load.Avg()
	if err == nil {
		metric = common.NewMetric("cpu.load1", cpuload.Load1)
		ret = append(ret, metric)

		metric = common.NewMetric("cpu.load5", cpuload.Load5)
		ret = append(ret, metric)
	}
	return
}

// MemMetric for memory
func MemMetric() (ret SPtr2Metric) {
	var metric *common.Metric

	memstat, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	percent := memstat.UsedPercent * float64(memstat.Total)
	metric = common.NewMetric("mem.percent", percent)
	ret = append(ret, metric)
	metric = common.NewMetric("mem.used", float64(memstat.Used))
	ret = append(ret, metric)
	return
}

// DiskMetric for /
func DiskMetric() (ret SPtr2Metric) {
	var metric *common.Metric

	root, err := disk.Usage("/")
	if err != nil {
		panic(err)
	}
	metric = common.NewMetric("disk.percent", root.UsedPercent)
	ret = append(ret, metric)
	metric = common.NewMetric("disk.usage", float64(root.Used))
	ret = append(ret, metric)

	return
}
