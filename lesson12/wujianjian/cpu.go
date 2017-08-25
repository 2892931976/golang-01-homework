package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	// CPU 信息
	cpus, err := cpu.Percent(time.Second, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cpu: %v\n", cpus)

	// load 信息
	loadavg, err := load.Avg()
	if err != nil {
		panic(err)
	}

	fmt.Printf("load:%v\n", loadavg)

	// mem 信息
	memstat, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	fmt.Printf("mem:%v\n", memstat.UsedPercent)

	// disk 信息
	diskinfo, err := disk.Usage("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("diskinfo:%v\n", diskinfo)

	// host 信息
	hostinfo, err := host.Info()
	if err != nil {
		panic(err)
	}
	fmt.Printf("hostinfo:%v\n", hostinfo)
}
