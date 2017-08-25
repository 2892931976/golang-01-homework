package main

import (
	"encoding/json"
	"fmt"
	"lesson12/monitor/common"
	"log"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {

	hostname, _ := os.Hostname()
	tag := []string{runtime.GOOS}
	for {
		cpus, err := cpu.Percent(time.Second, false)
		memstat, err := mem.VirtualMemory()
		if err != nil {
			panic(err)
		}
		metric := &common.Metric{
			Metric:    "cpu.usage",
			Endpoint:  hostname,
			Value:     cpus[0],
			Tag:       tag,
			Timestamp: time.Now().Unix(),
		}
		metric1 := &common.Metric{
			Metric:    "mem.usage",
			Endpoint:  hostname,
			Value:     memstat.UsedPercent,
			Tag:       tag,
			Timestamp: time.Now().Unix(),
		}
		buf, _ := json.Marshal(metric)
		buf1, _ := json.Marshal(metric1)
		fmt.Println(string(buf))
		fmt.Println(string(buf1))
		conn, err := net.Dial("tcp", "59.110.12.72:6000")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		conn.Write(buf)
		conn.Write([]byte("\n"))
		conn.Write(buf1)
		conn.Write([]byte("\n"))
		time.Sleep(time.Second * 5)
	}
}
