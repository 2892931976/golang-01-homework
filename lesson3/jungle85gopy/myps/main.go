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
	fmt.Println("pid\tprocess name")
	infos, err := f.Readdir(-1)
	for _, info := range infos {
		if _, err := strconv.Atoi(info.Name()); err != nil {
			continue
		} else if info.IsDir() {
			fname := "/proc/" + info.Name() + "/cmdline"
			buf, err := ioutil.ReadFile(fname)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\t%s \n", info.Name(), string(buf))
		}
	}
}
