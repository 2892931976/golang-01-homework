package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	s := Student{
		Id:   2,
		Name: "alice",
	}
	f, _ := os.Create("serialize.txt")
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("marshal error:%s", err)
	}
	fmt.Fprintln(f, string(buf))

	f.Close()
}
