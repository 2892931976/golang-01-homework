package main

import (
	"log"
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/xiaoran/monitor/common"
)

type Sched struct {
	ch chan *common.Metric
}

type MetricFunc func() []*common.Metric

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
			s.ch <- metric
			log.Print(metric)
		}
	}
}
