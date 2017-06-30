package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	count := make(map[string]int)                                                               //初始化一个map变量count,key为字符串，值为int
	buf, err := ioutil.ReadFile("/Users/yhzhao/Documents/workspace/Go/src/lesson4/mymap/a.txt") //读取文件内容
	if err != nil {
		fmt.Println(err)
		return
	}
	words := strings.Fields(string(buf)) //对读取的字符串，以空格为分割成一个数组
	for _, word := range words {         //循环单词数组
		count[word] += 1 //在count中查找key为单词的值，如果没找到值为0，对值进行加1
	}
	for k, v := range count { //循环count的key和值
		fmt.Printf("单词：%-20s出现次数：%d\n", k, v)
	}

}
