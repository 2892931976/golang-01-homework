package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("tock")
			return
		default:
			fmt.Println("eat someting")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
