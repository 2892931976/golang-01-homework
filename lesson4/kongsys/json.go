package main

import (
    "os"
    "encoding/json"
    "fmt"
    "log"
)

type Student struct {
    Id int
    Name string
}

func main() {
    s := Student{
        Id: 2,
        Name: "alice",
    }

    buf, err := json.Marshal(s)
    if err != nil {
        log.Fatalf("marshal error:%s", err) 
    }
    fmt.Println(string(buf))
    f, err := os.Create("std.db")
    if err != nil {
        log.Fatal("open db file error.") 
    }
    f.Write([]byte(buf))
    f.Write([]byte("\n"))
    f.Close()
}
