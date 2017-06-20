package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

func main() {
	var pids []int
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		n, err := strconv.Atoi(file.Name())
		if err != nil {
			continue
		}
		pids = append(pids, n)
	}
	sort.Ints(pids)

	fmt.Println("PID\tCMD")
	for _, pid := range pids {
		filename := fmt.Sprintf("%v/%v/%v", "/proc", pid, "cmdline")
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			continue
		}
		fmt.Printf("%v\t%v\n", pid, string(data))
	}
}
