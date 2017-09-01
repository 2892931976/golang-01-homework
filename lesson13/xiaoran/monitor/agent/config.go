package main

import "flag"

type UserScripConfig struct {
	Path string
	Step int
}

type SenderConfig struct {
	TransAddr     string `toml:"trans_addr"`
	FlushInterval int    `toml:"flush_interval"`
	MaxSleepTime  int    `toml:"max_sleep_time"`
}

type config struct {
	Sender     SenderConfig
	UserScript []UserScripConfig `toml:"user_script"`
}

var (
	configPath = flag.String("config", "config.toml", "config path")
	gcfg       config
)
