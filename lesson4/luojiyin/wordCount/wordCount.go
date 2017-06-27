package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	fs, err := ioutil.ReadFile("a.txt")
	if err != nil {
		return
	}
	temp := strings.Fields(string(fs))
	//fmt.Printf("Fields are:%q", temp)
	//for _, word := range temp {
	//	fmt.Println(word)
	//}
	//fmt.Println(countWord(temp))
	temp1 := countWord(temp)
	for k, v := range temp1 {
		fmt.Println(k, v)
	}
	pl := make(PairList, len(temp1))
	i := 0
	for k, v := range temp1 {
		pl[i] = Pair{k, v}
		i++
	}
	fmt.Println("-------------------------")
	fmt.Println("after sort")
	fmt.Println("--------------------------")
	sort.Sort(pl)
	for _, k := range pl {
		fmt.Printf("%s %d\n", k.Key, k.Value)
	}
}

func countWord(words []string) map[string]int {
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++

	}
	return wordCounts

}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
