package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f1, err := os.Open("/proc")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("PID\tCMD")
	infos, _ := f1.Readdir(-1)
	for _, info := range infos {
		_, err := strconv.Atoi(info.Name())
		if err != nil {
			continue
		}
		f2, err := os.Open("/proc/" + info.Name())
		if err != nil {
			continue
		}
		infos1, _ := f2.Readdir(-1)
		for _, info1 := range infos1 {
			if info1.Name() == "cmdline" {
				f3, err := os.Open("/proc/" + info.Name() + "/cmdline")
				if err == nil {
					r := bufio.NewReader(f3)
					for {
						line, err := r.ReadString('\n')
						fmt.Printf("%v\t%v\n", info.Name(), line)
						if err == io.EOF {
							break
						}
					}
				}
				f3.Close()
			}
		}
		f2.Close()
	}
	f1.Close()
}
