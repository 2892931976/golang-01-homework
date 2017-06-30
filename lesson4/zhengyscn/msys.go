package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

func Save(s []map[string]Student, filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range s {
		jsonByte, _ := json.Marshal(v)
		f.WriteString(string(jsonByte))
		f.WriteString("\n")
		fmt.Println(string(jsonByte))
	}
	f.Close()

}

func Load(filename string) []map[string]Student {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var s map[string]Student
	for _, info := range strings.Split(string(contents), "\n") {
		if len(info) != 0 {
			err := json.Unmarshal([]byte(info), &s)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	var sroom []map[string]Student
	for k, v := range s {
		var ss = make(map[string]Student)
		ss[k] = v
		sroom = append(sroom, ss)
	}
	return sroom
}

func List(sroom []map[string]Student) {
	fmt.Println("Id   Name")
	for _, room := range sroom {
		for _, stu := range room {
			fmt.Printf("%#v    %v\n", stu.Id, stu.Name)
		}
	}
}

func ValidExits(classroom []map[string]Student, name string) bool {
	fmt.Println(len(classroom))
	if len(classroom) == 0 {
		return false
	}
	for _, m := range classroom {
		for _, info := range m {
			fmt.Println(info.Name, name)
			if name == info.Name {
				fmt.Println("err")
				return true
			}
		}
	}
	return false

}

func main() {
	var cmd string
	var name string
	var id int
	f := bufio.NewReader(os.Stdin)
	var students map[string]Student
	var classroom []map[string]Student

	for {
		fmt.Printf("> ")
		line, _ := f.ReadString('\n')
		fmt.Sscan(line, &cmd, &name, &id)
		switch cmd {
		case "add":
			var stu Student
			stu.Id = id
			stu.Name = name
			students = make(map[string]Student)
			students[name] = stu

			if ValidExits(classroom, name) {
				fmt.Printf("name:%v already exists\n", name)
				continue
			}
			classroom = append(classroom, students)

		case "list":
			List(classroom)
		case "save":
			Save(classroom, name)
		case "load":
			classroom = Load(name)
		case "exit":
			fmt.Println("Bye!!!")
			break
		default:
			fmt.Println("alert choice ['add', 'list', 'save', 'load']")
		}
	}
}
