package main

import (

"encoding/json"
"fmt"
"log"
)
type Student struct {
   Name  string
   Id  int
}

func main() {

s := Student{
    Name: "zhangsan",
    Id: 1,
}

buf, err := json.Marshal(s)

if err != nil {

    log.Fatal(err)

}


fmt.Println(string(buf))
}



