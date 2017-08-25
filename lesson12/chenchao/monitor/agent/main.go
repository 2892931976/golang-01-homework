package main

import (
	"os"
	"time"
	"runtime"
	"github.com/shirou/gopsutil/cpu"
	"monitor/common"

	"flag"
	"github.com/shirou/gopsutil/load"
	"fmt"
)

var (
	transAddr = flag.String("trans", "59.110.12.72:6000", "transder address")
)

func NewMetric(metric string, value float64) *common.Metric {
	hostname, _ := os.Hostname()
	return &common.Metric{
		Metric: metric,
		Endpoint:hostname,
		Value:value,
		Tag: []string{runtime.GOOS},
		Timestamp: time.Now().Unix(),
	}
}

// 收集cpu需要监控的指标数据
// 这里返回的是一个metric的切片，目的就是为了存放一个监控需要的多个指标数据
func CpuMetric() []*common.Metric {

	var ret []*common.Metric
	cpus, err := cpu.Percent(time.Second, false)
	if err != nil{
		panic(err)
	}
	fmt.Println("cpu.usage", cpus[0])
	cpu_metric := NewMetric("cpu.usage", cpus[0])
	ret = append(ret, cpu_metric)

	cpu_load, err := load.Avg()
	fmt.Println("cpu.load")
	if err == nil{
		load_metric1 := NewMetric("cpu.load1", cpu_load.Load1)
		ret = append(ret, load_metric1)
		load_metric5 := NewMetric("cpu.load1", cpu_load.Load1)
		ret = append(ret, load_metric5)
	}

	return ret
}

func main() {
	flag.Parse()
	//创建网络模块
	send := NewSender(*transAddr)		// 初始化一个结构体  里面有addr，chan
	ch := send.Channel()				// 返回初始化中的chan 用于接受数据

	// 创建scheduler模块  返回的是schedule结构体，里面包含的是存放metric的chan
	sched:= NewSched(ch)
	fmt.Println("end new sched")
	// 将收集到的监控指标和监控时间 传送给调度器
	go sched.AddMetric(CpuMetric, time.Second *5)
	// mem  disk ...

	// 循环从chan取出数据 并发送给transfer 如果chan没有数据则会阻塞
	send.Start()

}
