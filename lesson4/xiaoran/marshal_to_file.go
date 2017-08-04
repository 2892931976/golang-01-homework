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

	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("marshal error: %s", err)
	}

	f, err := os.Create("123.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(f, string(buf))
	f.Close()
}
