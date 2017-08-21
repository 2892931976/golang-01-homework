/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import (
	"yinzhengjie/monitor/common"
	"os"
	"github.com/shirou/gopsutil/cpu"
	"time"
	"flag"
	"fmt"
)

var   (
	transAddr = flag.String("trans","59.110.12.72:6000","transfer address")
)


func main() {
	flag.Parse()  //解析命令行参数
	Sender := NewSender(*transAddr)

	go  Sender.Start()

	ch := Sender.Channel() //定义一个channel

	ticker := time.NewTicker(5*time.Second)  //定义一个定时器

	for range ticker.C{  //采集数据，每个5秒发送一次。
		hostname,_ := os.Hostname()
		cpus,err := cpu.Percent(time.Second,true)
		if err != nil {
			panic("获取CPU信息失败！")
		}

		metric := &common.Metric{
			Metric:"CPU.usage",
			Endpoint:hostname,
			Value:cpus[1],
			Timestamp:time.Now().Unix(),
		}
		fmt.Println(metric)
		ch <- metric  //往channle发送数据。
	}
}

