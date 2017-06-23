package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	dirs, err := ioutil.ReadDir("/proc")
	if err != nil {
		fmt.Println("read dir error.")
		os.Exit(1)
	}
	for _, d := range dirs {
		if d.IsDir() {
			if pid, err := strconv.Atoi(d.Name()); err == nil {
				pathname := fmt.Sprintf("/proc/%d/stat", pid)
				f, err := os.Open(pathname)
				if err != nil {
					continue
				}
				r := bufio.NewReader(f)
				line, _ := r.ReadString('\n')
				statslice := strings.Split(line, " ")
				app := statslice[1][1 : len(statslice[1])-1] // delete '(' and  ')', origin app like "(sshd)"
				fmt.Println(app, d.Name())
			}
		}
	}
}
