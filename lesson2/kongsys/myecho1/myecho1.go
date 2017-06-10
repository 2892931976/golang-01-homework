package main

import (
  "flag"
  "fmt"
  "strings"
)

var sep = flag.String("s", " ", "separator")
var nline = flag.Bool("n", false, "new line")

func main() {
  flag.Parse()
  if *nline {
    fmt.Println(strings.Join(flag.Args(), *sep))
  } else {
    fmt.Print(strings.Join(flag.Args(), *sep))
  }
}
