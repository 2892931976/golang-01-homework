package main

import (
	 "fmt"
	"time"



)

func main() {
	 tick := time.NewTicker(10000 * time.Millisecond).C
	 boom := time.After(10000 * time.Millisecond)

	 for {
		  select {
		      case <-tick:
			     fmt.Println("滴答...")
			     return
			  case <-boom:
			     fmt.Println("嘣！！！")
			     return
			  default:
			      fmt.Println("吃一口面")
			      time.Sleep(500 * time.Millisecond)


		  }


	 }


}
