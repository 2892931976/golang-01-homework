package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	var cmd string
	var name string
	var id int
	var line string
	f := bufio.NewReader(os.Stdin)
	var stuinfo = make(map[string]Student)
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			for _, v := range stuinfo {
				fmt.Println(v.Name, v.Id)
			}
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			stuinfo[name] = Student{
				Id:   id,
				Name: name,
			}
			fmt.Println("add done")
		case "save":
			w, err := json.Marshal(stuinfo)
			if err != nil {
				fmt.Println("save faile")
				return
			}
			ioutil.WriteFile("list.db", w, 0400)
			fmt.Println("save ok")
		case "load":
			f, err := ioutil.ReadFile("list.db")
			if err != nil {
				fmt.Println("load faile")
				return
			}
			json.Unmarshal(f, &stuinfo)
			fmt.Println("load ok")
		case "exit", "quit":
			fmt.Println("Bye Bye")
			os.Exit(0)
		default:
			fmt.Println("args: add user num | list | load | save | exit(quit)")
		}
	}
}
