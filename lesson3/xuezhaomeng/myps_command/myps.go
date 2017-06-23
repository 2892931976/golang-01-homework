package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func get_pid_name(s string) string {
	f, err := os.Open("/proc/" + s + "/cmdline")
	if err != nil {
		log.Fatal(err)

	}

	r := bufio.NewReader(f)
	var full string
	for {
		line, err := r.ReadString('\n')
		full = full + line
		if err == io.EOF {
			break
		}
	}

	f.Close()
	return full

}

func get_process_pid() {
	//arg := os.Args[1]
	arg := "/proc"
	f, err := os.Open(arg)

	if err != nil {
		log.Fatal(err)
	}

	infos, _ := f.Readdir(-1)
	for _, info := range infos {

		//fmt.Printf("%v  %d %s \n", info.IsDir(), info.Size(), info.Name())
		if info.IsDir() {
			n, err := strconv.Atoi(info.Name())
			if err == nil {
				//fmt.Println(n)

				pid_name := get_pid_name(strconv.Itoa(n)) //获取pid对应的name; n为int类型,转换为str
				if pid_name != "" {                       //判断是否为空,为空不输出
					fmt.Printf("%d\t%s\n", n, pid_name)
				}

			}

		}

	}

}

func main() {
	get_process_pid()
}
