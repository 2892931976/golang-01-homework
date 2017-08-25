package main

import (
	"flag"

	"github.com/golang-01-homework/lesson12/UFO-X/monitor/agent"
)

var (
	addr = flag.String("a", "59.110.12.72:6000", "stransfer ip:port")
)

func main() {
	// addr
	// agent type {addr metric}
	//trans type{addr metric}
	// (agent)addfunc(agentfunc time )
	flag.Parse()

	ag := agent.NewAgent()
	sd := agent.NewSched(*addr, ag.Ch)

	go ag.AddMertric(agent.Cpu_metric, 5)
	go ag.AddMertric(agent.Mem_metric, 5)
	go ag.AddMertric(agent.Disk_metric, 5)
	sd.Send()
}
