package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

func main() {
	f, err1 := os.Open("/proc")
	if err1 != nil {
		log.Fatal(err1)
	}
	infos, _ := f.Readdir(-1)
	for _, info := range infos {
		if (info.IsDir()){
		s := "/proc/" + info.Name()
		t,_ := os.Open(s)
		cmdlines, _ := t.Readdir(-1)
		for _ , cmdline := range cmdlines{
		   if (cmdline.Name() == "cmdline") {
			pspath := s + "/" + cmdline.Name()
			ps,_ := ioutil.ReadFile(pspath)
			if (string(ps) != ""){
				fmt.Printf("%s     %s\n",info.Name(),string(ps))
				}
			}
		}
	}
      }
}
