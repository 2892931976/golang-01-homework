package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	/*
		1. 粗暴方式：strings.Fields,Fields比Split高效
		2. 高效方式：
	*/

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	count := make(map[string]int)
	//Fields返回[]string切片
	worlds := strings.Fields(string(content))
	for _, world := range worlds {
		//fmt.Println(world)
		/* world存在于count
		   if 1 {
			//count里面对应的world的value加1
		   } else {
			//置count里面对应的world的value为初值1

		}
		*/
		if v, ok := count[world]; ok {
			v++
			count[world] = v
		} else {
			count[world] = 1
		}
	}

	for world, cnt := range count {
		fmt.Println(world, cnt)
	}

}
