package main

import (
	"fmt"
	"log"
	"os"

	"strconv"
)

func main() {
	f, err := os.Open("/proc")
	if err != nil {
		log.Fatal(err)
		return
	}
	infos, err := f.Readdir(-1)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, info := range infos {
		fmt.Println(info.IsDir(), info.Name(), info.Size())
		_, err := strconv.Atoi(info.Name())
		if err != nil {
			continue
		}

		if info.IsDir() {
			p_filename := "/proc/" + info.Name() + "/cmdline"
			fd, err := os.Open(p_filename)
			if err != nil {
				log.Fatal(err)
				continue
			}
			buf := make([]byte, 1024)
			fd.Read(buf)

			fmt.Printf("%s %v\n", info.Name(), string(buf))
		}
	}
}
