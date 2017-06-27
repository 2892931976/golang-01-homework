package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var rootdir string = os.Args[1] + "/"

const COMMAND = "/cmdline"

func main() {
	fmt.Println(rootdir, COMMAND)
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// -1 表示没有限制
	infos, _ := f.Readdir(-1)
	for _, info := range infos {
		if info.IsDir() {
			//fmt.Printf("%v %d %s\n", info.IsDir(), info.Size(), info.Name())
			pid, err := strconv.Atoi(info.Name())
			if err != nil {
				//fmt.Println("error: ", err)
			} else {
				//fmt.Println(pid)
				res, err := ioutil.ReadFile(rootdir + strconv.Itoa(pid) + COMMAND)
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Printf("%5d  %s\n", pid, res)
				}
			}

		}
	}
	f.Close()
}
