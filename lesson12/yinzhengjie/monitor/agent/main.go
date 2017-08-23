/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import (
	"flag"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"os"
	"runtime"
	"time"
	"yinzhengjie/monitor/common"
)

var (
	transAddr = flag.String("trans", "59.110.12.72:6000", "transfer address")
)

func NewMetric(metric string, value float64) *common.Metric { //构造出一个辅助函数，可以简化我们的代码。
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
	metric := NewMetric("cpu.usage", cpus[0]) //调用cpu的指标
	ret = append(ret, metric)

	cpuload, err := load.Avg()
	if err == nil {
		metric = NewMetric("cpu.load1", cpuload.Load1) //采集一分钟内cpu的Load指标，当然只会采集到linux的load，在windows上是没有load指标的。
		ret = append(ret, metric)
		metric = NewMetric("cpu.load5", cpuload.Load5) //采集五分钟内cpu的Load指标，当然只会采集到linux的load，在windows上是没有load指标的。
		ret = append(ret, metric)
		metric = NewMetric("cpu.load5", cpuload.Load15) //采集十五分钟内cpu的Load指标，当然只会采集到linux的load，在windows上是没有load指标的。
		ret = append(ret, metric)
	}
	return ret
}

func main() {
	flag.Parse() //解析命令行参数
	sender := NewSender(*transAddr)
	ch := sender.Channel()

	sched := NewSched(ch)                        //构造一个调度器
	go sched.AddMetric(CpuMetric, time.Second*2) //定义调度器的调度周期，表示两秒钟采集一次数据。
	// memory, time.Second * 3
	// disk, time.Minute
	sender.Start()
}
