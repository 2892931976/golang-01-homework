package main

import (
	"monitor/common"
	"time"
	"fmt"
)

// 因为要通过一个函数 来获取数据
// 定义一个func类型 返回的是含有Metric的切片，因为有的监控项会有很多指标
type MetricFunc func() []*common.Metric

type Sched struct {
	ch chan *common.Metric
}

func NewSched(ch chan *common.Metric) *Sched {
	fmt.Println("New Sched")
	s := &Sched{ch: ch,}
	return s
}

func (s *Sched) AddMetric(collector MetricFunc, step time.Duration) {
	// 接受收集数据的方法 和 时间间隔
	// 利用时间间隔 去调用收集数据方法
	// collector 方法返回的是一个数据切片， 可能包含多个metric
	// 将得到的数据 传给chan
	// 调用start方法从chan取出数据 发送给server
	fmt.Println("Add Metric")
	timer := time.NewTicker(step)
	fmt.Println("create timer", step)
	for range timer.C{
		fmt.Println("time.time")
		metrics := collector()
		fmt.Println("======", metrics)
		for _, m := range metrics{
			fmt.Println("mmmmmm", &m)
			s.ch <- m
			fmt.Println("send data to chan over")
		}
	}
}