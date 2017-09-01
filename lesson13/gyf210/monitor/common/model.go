package common

import (
	"os"
	"runtime"
	"time"
)

var hostname string

type Metric struct {
	Metric    string   `json:"metric"`
	EndPoint  string   `json:"endpoint"`
	Tag       []string `json:"tag"`
	Value     float64  `json:"value"`
	Timestamp int64    `json:"timestamp"`
}

func init() {
	hostname, _ = os.Hostname()
}

func NewMetric(metric string, value float64) *Metric {
	return &Metric{
		Metric:    metric,
		EndPoint:  hostname,
		Tag:       []string{runtime.GOOS},
		Value:     value,
		Timestamp: time.Now().Unix(),
	}
}
