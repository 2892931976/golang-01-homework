/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/


package main

import (
	"bufio"
	"os"
	"fmt"

	"strconv"
)

var   (
	s string
	line string
)
func main()  {
	f := bufio.NewReader(os.Stdin)
	num := []int{100,200,300,400,500,600,700,800}
	fmt.Printf("现有一些数字：·\033[32;1m%v\033[0m·\n",num)
	for {
		fmt.Print("请您想要反转下标的起始的位置>")
		line,_ = f.ReadString('\n')
		if len(line) == 1 {
			continue  //过滤掉空格；
		}
		fmt.Sscan(line,&s)
		if s == "stop" {
			break //定义停止程序的键值;
		}
		index,err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("对不起，您必须输入一个数字")
		}
		num1 := num[:index]
		num2 := num[index:]
		i := 0
		for {

			num2=append(num2, num1[i])
			i = i + 1
			if i >= len(num1) {
				break
			}
		}
		fmt.Printf("反转后的内容是·\033[31;1m%v\033[0m·\n",num2)
	}
}
