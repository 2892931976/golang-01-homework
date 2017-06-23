/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func catFile(pid string, fileName string) {
	var s string
	buf, err := ioutil.ReadFile(fileName) //读取文件内容，即pid/cmdline文件内容
	if err != nil {
		log.Fatal(err) //如果文件不存在或是权限异常就打印出来，并且会附带时间戳效果哟！
		return
	}
	s = string(buf) //将读取的内容变成字符串。
	if len(s) > 0 { //如果读取的cmdline文件的内容不为空，就打印其的PID（数字目录）和进程名称。
		fmt.Printf("进程的PID是：%v\t进程的名称是：%v\n", pid, s)
	} else {

		fmt.Println("我天，这个PID进程是属于系统的BASH的哟！")
	}
}

func main() {
	var fileName string
	if 1 == len(os.Args) {
		fmt.Println("对不起，您未输入的有效参数,请重新输入！\n")
		return
	}
	f, err := os.Open(os.Args[1]) //打开命令行位置参数下表为1的文件。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("真是抱歉，您输入的'%s'并没有cmdline文件,您可以输入‘/proc’目录试试!\n", os.Args[1])
	infos, _ := f.Readdir(-1) //获取目录下的所有文件和目录。
	for _, info := range infos {
		_, err := strconv.Atoi(info.Name())
		if info.IsDir() && err == nil { //判断是否为目录，并且转换成int类型时无报错
			fileName = os.Args[1] + info.Name() + "/cmdline" //拼接其绝对路径。
			catFile(info.Name(), fileName)                   //将数字的目录和其绝对路径传给函数catFile函数。
		}
	}
	f.Close()
}
