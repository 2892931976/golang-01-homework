package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func getcmdline(pid string, cmdlinepath string) {
	f, err := os.Open(cmdlinepath)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	f.Read(buf)
	fmt.Printf("%v\t%v\n", pid, string(buf))
	f.Close()
}

func main() {

	f, err := os.Open("/proc")
	if err != nil {
		log.Fatal(err)
	}

	infos, _ := f.Readdir(-1)

	for _, info := range infos {

		if info.IsDir() {
			_, err := strconv.Atoi(info.Name())
			if err == nil {
				_path := "/proc/" + info.Name() + "/cmdline"
				getcmdline(info.Name(), _path)
			}
		}
	}
	f.Close()
}
