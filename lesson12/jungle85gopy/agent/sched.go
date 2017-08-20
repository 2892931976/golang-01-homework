package main

import (
	"sync"
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/jungle85gopy/common"
)

// MetricFunc metric func who return ptr slice of metric
type MetricFunc func() SPtr2Metric

// Sched struct based on chan
type Sched struct {
	ch chan *common.Metric
	wg sync.WaitGroup
}

// NewSched new Sched based on chan
func NewSched(ch chan *common.Metric) *Sched {
	return &Sched{
		ch: ch,
	}
}

// AddMetric add metric to Sched
func (s *Sched) AddMetric(collector MetricFunc, step time.Duration) {
	s.wg.Add(1)
	go func() {
		ticker := time.NewTicker(step)
		for range ticker.C {
			debugInfo("## ticker hit at" + time.Now().String())
			for _, metric := range collector() {
				debugInfo(metric)
				s.ch <- metric
			}
		}
	}()
}

// Wait to wait all routine.
func (s *Sched) Wait() {
	s.wg.Wait()
}
