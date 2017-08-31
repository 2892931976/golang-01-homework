package main

import (
	"flag"
)

type UserScriptConfig struct {
	Path string
	Step int
}

type SenderConfig struct {
	TransAddr     string `toml:"trans_addr"`
	FlushInterval int    `toml:"flush_interval"`
	MaxSleepTime  int    `toml:"max_sleep_time"`
	BufferLen     int    `toml:"buffer_len"`
	Debug         bool   `toml:"debug"`
}

type config struct {
	Sender     SenderConfig
	UserScript []UserScriptConfig `toml:"user_script"`
}

var (
	configPath = flag.String("config", "../config/agent.toml", "config path")
	gcfg       config
)
