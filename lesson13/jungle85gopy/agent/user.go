package main

import (
	"bufio"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/51reboot/golang-01-homework/lesson13/jungle85gopy/common"
)

// UserMetric for user metrics
func UserMetric() (ret SPtr2Metric) {
	userMetric, err := getUserMetrics("./userDefined/user.py")
	if err != nil {
		log.Println(err)
	}
	return userMetric
}

// NewUserMetric new user metric func
func NewUserMetric(cmdstr string) MetricFunc {
	return func() SPtr2Metric {
		metrics, err := getUserMetrics(cmdstr)
		if err != nil {
			log.Print("get user metric err:", err)
			return []*common.Metric{}
		}
		return metrics
	}
}

// 构造命令；获取标准输出；按行解析；获取key/value；包装common.Metric
func getUserMetrics(srcCmd string) (ret SPtr2Metric, err error) {
	cmd := exec.Command("bash", "-c", srcCmd)
	out, _ := cmd.StdoutPipe()
	if er := cmd.Start(); er != nil {
		log.Print("command run err:", er)
		return nil, er
	}
	f := bufio.NewReader(out)
	for {
		line, er := f.ReadString('\n')
		if er != nil {
			break
		}
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)
		value, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			log.Println(err)
			continue
		}
		metric := common.NewMetric(fields[0], value)
		ret = append(ret, metric)
	}
	// Wait waitfor the cmd to exit,
	// and release any resource associated with the cmd
	cmd.Wait()
	return
}
