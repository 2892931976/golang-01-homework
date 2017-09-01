package main

import (
	"bytes"
	"github.com/gyf210/monitor/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func CollectCpuMetric() []*common.Metric {
	var ret []*common.Metric
	cpus, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Print(err)
	} else {
		metric := common.NewMetric("cpu.usage", cpus[0])
		ret = append(ret, metric)
	}

	loadstat, err := load.Avg()
	if err != nil {
		log.Print(err)
	} else {
		metric := common.NewMetric("cpu.load1", loadstat.Load1)
		ret = append(ret, metric)
		metric = common.NewMetric("cpu.load5", loadstat.Load5)
		ret = append(ret, metric)
		metric = common.NewMetric("cpu.load15", loadstat.Load15)
		ret = append(ret, metric)
	}
	return ret
}

func CollectMemMetric() []*common.Metric {
	var ret []*common.Metric
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Println(err)
	} else {
		metric := common.NewMetric("mem.usage", m.UsedPercent)
		ret = append(ret, metric)
	}
	return ret
}

func CollectScriptMetric(cmd string) MetricFunc {
	return func() []*common.Metric {
		metrics, err := GetScriptMetric(cmd)
		if err != nil {
			log.Println(err)
			return []*common.Metric{}
		}
		return metrics
	}
}

func GetScriptMetric(cmdline string) ([]*common.Metric, error) {
	var ret []*common.Metric
	cmd := exec.Command("bash", "-c", cmdline)
	result, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(result)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		key, value := fields[0], fields[1]
		n, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		metric := common.NewMetric(key, n)
		ret = append(ret, metric)
	}
	return ret, nil
}
