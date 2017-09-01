package main

import (
	"flag"
)

type BaseConfig struct {
	CpuStep int `toml:"cpu_step"`
	MemStep int `toml:"mem_step"`
}

type UserScriptConfig struct {
	Path string `toml:"path"`
	Step int    `toml:"step"`
}

type SenderConfig struct {
	TransAddr     string `toml:"trans_addr"`
	FlushInterval int    `toml:"flush_interval"`
	MaxSleepTime  int    `toml:"max_sleep_time"`
}

type GlobalConfig struct {
	Base       BaseConfig         `toml:"base"`
	Sender     SenderConfig       `toml:"sender"`
	UserScript []UserScriptConfig `toml:"user_script"`
}

var (
	configPath = flag.String("config", "config.toml", "config path")
	gconfig    GlobalConfig
)
