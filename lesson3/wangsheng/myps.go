package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const (
	procdir     = "/proc/"
	proccommand = "/cmdline"
)

func main() {
	f, err := os.Open(procdir)
	//fmt.Println(f)  print &{0xc82000e1e0}
	if err != nil {
		log.Fatal(err)
	}

	files, _ := f.Readdir(-1)
	//fmt.Println(files)
	//[0xc82006c000 0xc82006c0d0 0xc82006c1a0 0xc82006c270 0xc82006c340 0xc82006c410 0xc82006c4e0

	for _, file := range files {
		//fmt.Println(file)
		//&{30734 0 2147484013 {63632590984 45532920 0x589800} {3 1332662 9 16749 0 0 0 0 0 1024 0 {1496994184 45532920} {1496994184 45532920} {1496994184 45532920} [0 0 0]}}
		if file.IsDir() {
			filename, err := strconv.Atoi(file.Name())
			//fmt.Println(filename)
			//0 1 2 3 7 8 9  10  11 12
			if err == nil {
				filebuf, _ := ioutil.ReadFile(procdir + strconv.Itoa(filename) + proccommand)
				fmt.Println(filename, string(filebuf))
			}
		}
	}

}
