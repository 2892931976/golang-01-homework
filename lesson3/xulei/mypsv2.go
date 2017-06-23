package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("/proc")
	if err != nil {
		log.Fatal(err)
	}
	infos, _ := f.Readdir(-1)
	for _, info := range infos {
		// fmt.Printf("%v %d %s\n", info.IsDir(), info.Size(), info.Name())
		if info.IsDir() {
			// fmt.Printf("%s\n", info.Name())
			_, err := strconv.Atoi(info.Name())
			if err == nil {
				pidpath := info.Name()
				pidfile := "/proc" + "/" + pidpath + "/" + "cmdline"
				// fmt.Println(pidfile)
				//fmt.Println("--------------------")
				//data, err := ioutil.ReadFile(pidfile)
				data, err := ioutil.ReadFile(pidfile)
				if err == nil {
					fmt.Printf("%s %s\n", pidpath, string(data))
				}
			}

		} else {
			continue
		}

	}
	f.Close()
}
