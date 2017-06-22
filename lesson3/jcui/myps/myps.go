package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"io/ioutil"
)

func psfile(pid, filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	s := string(buf)

	if len(s) != 0 {
		fmt.Printf("%s \t %s \n", pid, s)
	}
}

func main() {
	f, err := os.Open("/proc/")
	if err != nil {
		log.Fatal(err)
	}
	infos, err := f.Readdir(-1)
	//names, err := f.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, args := range infos {
		r, _ := regexp.MatchString(`[0-9]`, args.Name())
		if r && args.IsDir() {
			psfile(args.Name(), "/proc/"+args.Name()+"/cmdline")
			//fmt.Println(args.Size())
			//fmt.Println(args.ModTime())
			//fmt.Println(args.Mode())
		}
	}

}
