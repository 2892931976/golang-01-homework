package main

import (
    "log"
    "bufio"
    "encoding/json"
    "os"
    "fmt"
)

type Student struct {
    Id int
    Name string
}
func main() {
    f, err := os.Open("std.db")
    if err != nil {
        log.Fatalf("open db file error:%s", err) 
    }
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        str := scanner.Text()
        var s Student
        err := json.Unmarshal([]byte(str), &s)
        if err != nil {
            log.Fatalf("Unmarshal error:%s", err)
        }
        fmt.Println(s) 
    }

}
