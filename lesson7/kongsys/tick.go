package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(time.Second)
	cnt := 0
	for _ = range timer.C {
		cnt++
		if cnt > 10 {
			timer.Stop() // only stop send data to channel
			return
		}
		fmt.Println("hello")
	}
}
