package main

import (
	"fmt"
	"strings"
)

func main() {
	var word = "Reboot Goloag"
	fmt.Println(word)
	fmt.Println(len(strings.Fields(word)))
	tmp := strings.Fields(word)
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	fmt.Println(tmp)

}
