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
		log.Fatal(err)
		return
	}
	s = string(buf)
	if len(s) == 0 {
		//如果cmdline为空则不打印
	} else {
		fmt.Printf("%v\t%v\n", pid, s)
	}
}

func main() {
	var fileName string
	f, err := os.Open("/proc/")
	if err != nil {
		log.Fatal(err)
	}
	infos, _ := f.Readdir(-1)
	for _, info := range infos {
		_, err := strconv.Atoi(info.Name())
		if info.IsDir() && err == nil {
			fileName = "/proc/" + info.Name() + "/cmdline"
			catFile(info.Name(), fileName)
		}
	}
	f.Close()
}
