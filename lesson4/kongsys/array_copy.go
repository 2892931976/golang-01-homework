package main

import "fmt"
import "unsafe"

func main() {
    a1 := [3]int{1, 2, 3}
    var a2 [3]int
    a2 = a1
    fmt.Println(&a1[0], &a2[0])
    fmt.Println(&a1[1], &a2[1])
    fmt.Println(&a1[2], &a2[2])
    fmt.Println(&a1, &a2)
    fmt.Println(unsafe.Sizeof(a1))
    fmt.Println(unsafe.Sizeof(a2))
    fmt.Printf("%x\n", 255)
}
