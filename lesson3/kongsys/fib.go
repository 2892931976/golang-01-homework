package main

import "fmt"

func main() {
  var sum int = 0
  var i1, i2 int = 0, 1
  for i2 < 100 {
    i1, i2 = i2, i1 + i2 
    sum += i1
  }
  fmt.Println(sum)

}
