package main

import (
	"bufio"
	"fmt"
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

	var username string
	var password string
	f := bufio.NewReader(os.Stdin)

	//var students map[string]Student
	
	fmt.Print("欢迎登入学生管理系统\n")
	fmt.Sscan(line, &cmd, &name, &id)
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Printf("binggan 01\njack 02\n")
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			fmt.Printf("add done.\n")
		}
	}
}
