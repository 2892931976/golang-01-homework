package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Monitor struct {
	counter int
}

func (m *Monitor) Run() {
	for {
		time.Sleep(time.Second)
		m.counter++
	}
}

func (m *Monitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "count:%d\n", m.counter)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func main() {
	var m Monitor
	http.HandleFunc("/", handle)
	http.Handle("/monitor", &m)
	go m.Run()
	log.Fatal(http.ListenAndServe(":19090", nil))
}
