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
	"time"
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"runtime"
)

var   (
	transAddr = flag.String("trans","59.110.12.72:6000","transfer address")
)

func NewMetraic(metric string, value float64) *common.Metric  {  //定义一个辅助函数，对其metric进行封装。
	hostname,_ := os.Hostname()
	return &common.Metric{
		Metric:metric,
		Endpoint:hostname,
		Value:value,
		Tag:[]string{runtime.GOOS},
		Timestamp:time.Now().Unix(),
	}
}


func main() {
	flag.Parse()  //解析命令行参数
	
	Sender := NewSender(*transAddr)
	go  Sender.Start()

	ch := Sender.Channel() //定义一个channel

	ticker := time.NewTicker(5*time.Second)  //定义一个定时器

	for range ticker.C{  //采集数据，每个5秒发送一次。
		hostname,_ := os.Hostname()
		diskstat,err := disk.Usage("H:\\Golang进阶之路")
		if err != nil {
			panic(nil)
		}
		metric := NewMetraic("disk.used",diskstat.UsedPercent)
		fmt.Println(metric)
		ch <- metric  //往channle发送数据。
	}
}

