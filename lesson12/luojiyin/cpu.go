package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type Metric struct {
	Metric    string   `json:"metric"`
	Endpoint  string   `json:"endpoint"`
	Tag       []string `json:"Tag"`
	Value     float64  `json:"value"`
	Timestamp int64    `json:"timestamp"`
}

func main() {
	//cpus, err := cpu.Percent(time.Second, false)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(cpus[0])

	//loadavg, err := load.Avg()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(loadavg)

	conn, err := net.Dial("tcp", "59.110.12.72:6000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println(err)
		}
		cpus, err := cpu.Percent(time.Second, false)
		if err != nil {
			fmt.Println(err)
		}

		metric := &Metric{
			Metric:    "cpu.usage",
			Endpoint:  hostname,
			Value:     cpus[0],
			Tag:       []string{runtime.GOOS},
			Timestamp: time.Now().Unix(),
		}
		buf, _ := json.Marshal(metric)
		fmt.Println(string(buf))
		fmt.Fprintf(conn, "%s\n", buf)

	}
}
