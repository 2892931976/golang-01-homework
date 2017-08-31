package main

import (
	"flag"
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

func main() {
	flag.Parse()

	_, err := toml.DecodeFile(*configPath, &gcfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", gcfg)

	sender := NewSender(gcfg.Sender)
	go sender.Start()

	ch := sender.Channel()
	sched := NewSched(ch)
	sched.AddMetric(CPUMetric, time.Second*5)
	sched.AddMetric(MemMetric, time.Second*3)
	sched.AddMetric(DiskMetric, time.Second*15)
	// user defined metric
	for _, ucfg := range gcfg.UserScript {
		sched.AddMetric(NewUserMetric(ucfg.Path), time.Second*time.Duration(ucfg.Step))
	}

	sched.Wait()
}

// debugInfo output debug info
func debugInfo(info interface{}) {
	if gcfg.Sender.Debug {
		log.Printf("%v", info)
	}
}
