package main

import (
	"learning/mygo_local/lesson12/monitor/common"
	"log"
	"time"
)

type MetricFunc func() []*common.Metric

type Sched struct {
	ch chan *common.Metric
}

func NewSched(ch chan *common.Metric) *Sched {
	return nil
}

func (s *Sched) AddMetric(collecter MetricFunc, step time.Duration) {
	ticker := time.NewTicker(step)
	for range ticker.C {
		metrics := collecter()
		for _, metric := range metrics {
			log.Print(metric)

		}
	}
}
