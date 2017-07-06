package main

import "fmt"

func main() {
    s := make(map[string]int)

    s["ls"]=1

    _, ok := s["df"]
    fmt.Printf("%V %T\n", ok, ok)
    _, ok = s["ls"]
    fmt.Printf("%V %T\n", ok, ok)
}
