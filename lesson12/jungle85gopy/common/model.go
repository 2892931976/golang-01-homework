package common

import (
	"os"
	"runtime"
	"time"
)

// Metric struct
type Metric struct {
	Metric    string   `json:"metric"`
	Endpoint  string   `json:"endpoint"`
	Tag       []string `json:"tag"`
	Value     float64  `json:"value"`
	Timestamp int64    `json:"timestamp"`
}

// NewMetric new a Metric
func NewMetric(metric string, value float64) *Metric {
	hostname, _ := os.Hostname()

	return &Metric{
		Metric:    metric,
		Endpoint:  hostname,
		Value:     value,
		Tag:       []string{runtime.GOOS},
		Timestamp: time.Now().Unix(),
	}
}
