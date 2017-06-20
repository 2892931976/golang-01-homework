//
// ps.go
// Copyright (C) 2017 root <root@localhost.localdomain>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"io/ioutil"
	//"os"
	//"strconv"
	"regexp"
)

func main() {
	Dir := "/proc"
	fs, err := ioutil.ReadDir(Dir)
	if err != nil {
		return
	}

	size := len(fs)
	if size == 0 {
		return
	}

	result := make([]string, 0, size)
	for i := 0; i < size; i++ {
		if fs[i].IsDir() {
			name := fs[i].Name()
			var validPID = regexp.MustCompile(`^[0-9]+$`)
			if validPID.MatchString(name) {
				//fmt.Println(name)
				cmdFile := Dir + "/" + name + "/cmdline"
				cmdContext, e := ioutil.ReadFile(cmdFile)
				if e != nil {
					continue
				}
				if len(cmdContext) > 0 {
					fmt.Println(name, string(cmdContext))
					result = append(result, name)
				}
			}
		}
	}
	//fmt.Println(result)
}
