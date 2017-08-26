package main

import (
	"flag"
	"log"
	"time"
)

var (
	transAddr = flag.String("trans", ":6000", "transfer address")
	debug     = flag.Bool("debug", false, "debug data output")
)

func main() {
	flag.Parse()

	sender := NewSender(*transAddr)
	go sender.Start()

	ch := sender.Channel()
	sched := NewSched(ch)
	sched.AddMetric(CPUMetric, time.Second*5)
	sched.AddMetric(MemMetric, time.Second*3)
	sched.AddMetric(DiskMetric, time.Second*15)
	sched.Wait()
}

// debugInfo output debug info
func debugInfo(info interface{}) {
	if *debug {
		log.Printf("%v", info)
	}
}
