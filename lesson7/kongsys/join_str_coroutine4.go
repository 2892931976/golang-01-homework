package main

import "fmt"

func join_str(s []string, ch chan string) {
  var temp string
  for _, v := range s {
    temp = fmt.Sprintf("%s %s", temp, v)
  }
  ch <- temp
}

func main() {
  s := []string{"hello", "golang", "c++", "world"}

  c := make(chan string)
  go join_str(s[:len(s)/2], c) 
  go join_str(s[len(s)/2:], c) 
  x, y := <-c, <-c

  fmt.Printf("%s %s", x, y)
}
