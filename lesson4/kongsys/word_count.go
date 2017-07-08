package main

import (
    "fmt"
    _ "strings"
    "os"
    "log"
    "bufio"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("not enough parament.")
        return
    }
    f, err := os.Open(os.Args[1])
    defer f.Close()
    if err != nil {
        log.Fatal(err) 
    }
    word_count := make(map[string]uint)
    scanner := bufio.NewScanner(f)
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
 			if data[i] == ',' || data[i] == '.'  || data[i] == ' ' || data[i] == '?'{
				word_count[string(data[:i])]++
				return i + 1, data[:i], nil
			}
		}
		return 0, data, bufio.ErrFinalToken
	}
    scanner.Split(onComma)
    for scanner.Scan() {
    }
	for k, v := range word_count {
		fmt.Printf("%s : %d\n", k, v)
	}
}
