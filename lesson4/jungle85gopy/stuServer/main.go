package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

var students = make(map[string]Student)

func main() {
	var line, cmd, name string
	var id int
	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		line = strings.Replace(line, "\n", "", 1)
		fmt.Sscan(line, &cmd, &name, &id)
		if cmd == "exit" {
			break
		}
		// fmt.Println("input:", line, "len:", len(line))
		switch cmd {
		case "add":
			addStu(name, id)
		case "list":
			listStu()
		case "load":
			loadStu(name)
		case "save":
			saveStu(name)
		default:
			usage()
		}
	}
}

func saveStu(name string) {
	fmt.Println("params:", name)
}

func loadStu(name string) {
	fmt.Println("params:", name)
}

func addStu(name string, id int) {
	// fmt.Println("-- params:", name, id)
	if _, ok := students[name]; ok {
		fmt.Println("duplicated name:", name)
		return
	}
	students[name] = Student{Id: id, Name: name}
}

func listStu() {
	if len(students) == 0 {
		return
	}
	fmt.Println("Id\tName:")
	for _, val := range students {
		fmt.Printf("%d\t%s\n", val.Id, val.Name)
	}
}

func usage() {
	fmt.Println("cli usage:")
	fmt.Println("    add name id -- add student info")
	fmt.Println("    list \t-- list student info")
	fmt.Println("    load file \t-- load student from file")
	fmt.Println("    save file \t-- save student info file")
	fmt.Println("    exit \t-- exit the cli")
}
