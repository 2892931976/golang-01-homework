package main

import (
	"fmt"
	"net/http"
)

var (
	// just for check Counter is a http handler
	_ http.Handler = &Counter{}
)

// Counter struct
type Counter struct {
	h     http.Handler
	count map[string]int
}

// NewCounter new a counter
func NewCounter(h http.Handler) *Counter {
	return &Counter{
		h:     h,
		count: make(map[string]int),
	}
}

// GetCounter return all counter of path
func (c *Counter) GetCounter(w http.ResponseWriter, r *http.Request) {
	for path, count := range c.count {
		fmt.Fprintf(w, "%s\t%d\n", path, count)
	}
}

func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.count[r.URL.Path]++
	c.h.ServeHTTP(w, r)
}
