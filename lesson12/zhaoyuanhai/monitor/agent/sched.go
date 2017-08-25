package main

import (
	"lesson12/monitor/common"
	"time"
)

type MetricFunc func() []*common.Metric

type Sched struct {
	ch chan *common.Metric
}

func NewSched(ch chan *common.Metric) *Sched {
	return &Sched{
		ch: ch,
	}
}

func (s *Sched) AddMetric(collecter MetricFunc, step time.Duration) {
	ticker := time.NewTicker(step)
	for range ticker.C {
		metrics := collecter()
		for _, metric := range metrics {
			// log.Print("Sched Metric:", metric)
			s.ch <- metric
		}
	}
}
