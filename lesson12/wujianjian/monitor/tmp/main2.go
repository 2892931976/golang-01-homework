package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/wujianjian/monitor/common"
	"github.com/shirou/gopsutil/cpu"
)

var (
	transAddr = flag.String("trans", "59.110.12.72:6000", "transfer address")
)

func main() {
	flag.Parse()
	sender := NewSender(*transAddr)

	go sender.Start()
	ch := sender.Channel()

	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		hostname, _ := os.Hostname()
		cpus, err := cpu.Percent(time.Second, false)
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
		fmt.Println(metric)
		ch <- metric
	}
}
