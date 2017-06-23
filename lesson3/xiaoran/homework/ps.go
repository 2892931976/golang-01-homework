package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	var cmdlinename string
	f, err := os.Open("/proc")
	if err != nil {
		log.Fatal(err)

	}
	infos, _ := f.Readdir(-1)
	for _, info := range infos {
		_, err := strconv.Atoi(info.Name())
		if info.IsDir() == true && err == nil {
			cmdlinename = "/proc" + "/" + info.Name() + "/cmdline"
			buf, _ := ioutil.ReadFile(cmdlinename)
			fmt.Println(info.Name(), string(buf))

		}
	}

	f.Close()
}
