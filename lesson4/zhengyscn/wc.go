package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func Wc(d []byte) map[string]int {
	words := string(d)
	count := make(map[string]int)
	for _, word := range strings.Split(words, " ") {
		if _, ok := count[word]; ok {
			count[word] += 1
		} else {
			count[word] = 1
		}
	}
	return count
}

func Sort_slice(s []string, m map[string]int) {
	sort.Strings(s)
	for _, v := range s {
		fmt.Printf("%s -> %d\n", v, m[v])
	}
}

func main() {
	contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(err)
		return
	}
	mmwc := Wc(contents)
	s := []string{}
	for k := range mmwc {
		s = append(s, k)
	}
	Sort_slice(s, mmwc)
}
