package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//借助os包获取命令行参数
	var sum int
	for i := 1; i < len(os.Args); i++ {
		if 2 <= len(os.Args) && len(os.Args) <= 3 {
			n, err := strconv.Atoi(os.Args[i])
			if err != nil {
				fmt.Println("Error: ", err)
			}
			sum += n
		}
	}
	fmt.Println(sum)
}
