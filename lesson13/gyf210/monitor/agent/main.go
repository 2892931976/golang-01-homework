package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

func main() {
	flag.Parse()
	_, err := toml.DecodeFile(*configPath, &gconfig)
	if err != nil {
		log.Fatal(err)
	}
	sender := NewSender(gconfig.Sender.TransAddr)
	ch := sender.Channel()
	scheder := NewSched(ch)
	scheder.AddMetric(CollectMemMetric, time.Duration(gconfig.Base.MemStep)*time.Second)
	scheder.AddMetric(CollectCpuMetric, time.Duration(gconfig.Base.CpuStep)*time.Second)
	for _, cfg := range gconfig.UserScript {
		scheder.AddMetric(CollectScriptMetric(cfg.Path), time.Duration(cfg.Step)*time.Second)
	}
	err = sender.Start()
	if err != nil {
		log.Fatal(err)
	}
}
