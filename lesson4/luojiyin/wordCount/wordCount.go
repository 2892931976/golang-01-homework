package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fs, err := ioutil.ReadFile("a.txt")
	if err != nil {
		return
	}
	temp := strings.Fields(string(fs))
	fmt.Printf("Fields are:%q", temp)
	for _, word := range temp {
		fmt.Println(word)
	}
	fmt.Println(countWord(temp))
}

func countWord(words []string) map[string]int {
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++

	}
	return wordCounts

}
