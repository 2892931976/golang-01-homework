package agent

import (
	"os"
	"runtime"
	"time"

	"github.com/golang-01-homework/lesson12/UFO-X/monitor/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type Agent struct {
	Ch chan *common.Metric
}
type MetricFunc func() []*common.Metric

func NewAgent() Agent {

	return Agent{
		Ch: make(chan *common.Metric, 1000),
	}
}
func NewMetric(met string, val float64) *common.Metric {
	return &common.Metric{
		Metric:    met,
		Tag:       []string{runtime.GOOS},
		Value:     val,
		Timestamp: time.Now().Unix(),
	}
}

func (a *Agent) AddMertric(collect MetricFunc, t time.Duration) {
	tick := time.NewTicker(t * time.Second)
	host_name, _ := os.Hostname()
	for range tick.C { //v := range tick.C {
		// log.Println(v)
		mes := collect()
		for _, v := range mes {
			v.Endpoint = host_name
			a.Ch <- v
		}
	}
	tick.Stop()
}
func Cpu_metric() []*common.Metric {

	cpus, _ := cpu.Percent(time.Second, false)
	met := NewMetric("cpu.usage", cpus[0])
	var res []*common.Metric
	res = append(res, met)
	return res

}
func Mem_metric() []*common.Metric {
	mems, _ := mem.VirtualMemory()
	met := NewMetric("mem.usage", float64(mems.Free))
	var res []*common.Metric
	res = append(res, met)
	return res
}
func Disk_metric() []*common.Metric {
	disks, _ := disk.Usage(".")
	met := NewMetric("disk.usage", float64(disks.Free))
	var res []*common.Metric
	res = append(res, met)
	return res
}

// func net_metric() []*common.Metric {

// }
