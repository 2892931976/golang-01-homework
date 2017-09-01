package main

import (
	"github.com/gyf210/monitor/common"
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
	go func() {
		ticker := time.NewTicker(step)
		for range ticker.C {
			for _, metric := range collecter() {
				if metric != nil {
					s.ch <- metric
				}
			}
		}
	}()
}
