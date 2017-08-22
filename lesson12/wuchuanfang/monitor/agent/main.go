package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	//"github.com/51reboot/golang-01-homework/lesson12/wuchuanfang/monitor/common"
	"github.com/homework/monitor/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

var (
	addr = flag.String("trans", "59.110.12.72:6000", "transfer server")
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
	return ret
}

func MemMetric() []*common.Metric {
	var ret []*common.Metric
	memstat, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	metric := NewMetric("mem.usage", memstat.UsedPercent)

	ret = append(ret, metric)
	return ret
}

func DiskMetric() []*common.Metric {
	var ret []*common.Metric
	diskstat, err := disk.Usage("D:/")
	if err != nil {
		panic(err)
	}
	metric := NewMetric("disk.usage", diskstat.UsedPercent)

	ret = append(ret, metric)
	return ret
}

func main() {
	/*
		addr := "59.110.12.72:6000"
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Println(err)
		}
		defer conn.Close()

		for {
			hostname, _ := os.Hostname()
			cpus, err := cpu.Percent(time.Second, false)
			if err != nil {
				panic(err)
			}
			memstat, err := mem.VirtualMemory()
			if err != nil {
				panic(err)
			}

			metric := &common.Metric{
				Metric:    "cpu.usage",
				Endpoint:  hostname,
				Value:     cpus[0],
				Tag:       []string{runtime.GOOS},
				Timestamp: time.Now().Unix(),
			}

			memusage := &common.Metric{
				Metric:    "mem.usage",
				Endpoint:  hostname,
				Value:     memstat.UsedPercent,
				Tag:       []string{runtime.GOOS},
				Timestamp: time.Now().Unix(),
			}
			bufcpu, _ := json.Marshal(metric)
			bufmem, _ := json.Marshal(memusage)
			fmt.Println(string(bufcpu))
			fmt.Println(string(bufmem))

			conn.Write(bufcpu)
			conn.Write([]byte("\n"))
			conn.Write(bufmem)
			conn.Write([]byte("\n"))
			time.Sleep(5 * time.Second)
	*/

	//addr := "59.110.12.72:6000"

	/*
		flag.Parse()
		sender := NewSender(*addr)

		go sender.Start()
		ch := sender.Channel()

		ticker := time.NewTicker(time.Second * 5)
		for range ticker.C {

			hostname, _ := os.Hostname()
			cpus, err := cpu.Percent(time.Second, false)
			if err != nil {
				panic(err)
			}
			memstat, err := mem.VirtualMemory()
			if err != nil {
				panic(err)
			}

			metric := &common.Metric{
				Metric:    "cpu.usage",
				Endpoint:  hostname,
				Value:     cpus[0],
				Tag:       []string{runtime.GOOS},
				Timestamp: time.Now().Unix(),
			}

			memusage := &common.Metric{
				Metric:    "mem.usage",
				Endpoint:  hostname,
				Value:     memstat.UsedPercent,
				Tag:       []string{runtime.GOOS},
				Timestamp: time.Now().Unix(),
			}

			ch <- metric
			ch <- memusage

		}
	*/
	flag.Parse()
	fmt.Println(*addr)
	sender := NewSender(*addr)
	ch := sender.Channel()

	sched := NewSched(ch)

	go sched.AddMetric(CpuMetric, time.Second*5)
	go sched.AddMetric(MemMetric, time.Second*10)
	//fmt.Println(sender.ch)
	sender.Start()
}
