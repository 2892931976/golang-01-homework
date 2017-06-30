/*
完成学生信息管理系统

实现如下4个指令

1. add name id，添加一个学生的信息，如果name有重复，报错
2. list, 列出所有的学生信息
3. save filename，保存所有的学生信息到filename指定的文件中
4. load filename, 从filename指定的文件中加载学生信息

Code Example:

...
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type StudentSet struct {
	M map[string]*Student `json:"data"`
}

// NewStudentSets initial a new studentsets
func NewStudentSets() *StudentSet {
	return &StudentSet{M: make(map[string]*Student, 0)}
}

// Add student info
func (s *StudentSet) Add(id int, name string) (err error) {
	_, ok := s.M[name]
	if ok {
		err = fmt.Errorf("student %s already exists.", name)
		return
	}

	s.M[name] = &Student{Id: id, Name: name}

	return
}

// list all students info as a string
func (s *StudentSet) list() string {
	var str string = "Id\t\tName\n"
	for k, v := range s.M {
		str += fmt.Sprintf("%d\t\t%s\n", v.Id, k)
	}
	return str
}

// Remove student that specified
func (s *StudentSet) Remove(name string) {
	_, ok := s.M[name]
	if ok {
		delete(s.M, name)
	}
}

// Clear all students info
func (s *StudentSet) Clear() {
	s.M = make(map[string]*Student, 0)
}

// String implements String method
func (s *StudentSet) String() string {
	return s.list()
}

// Dump students info to the file that specified
func Dump(fileName string, stu *StudentSet) (err error) {
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fd.Close()

	bs, err := json.Marshal(stu)
	if err != nil {
		return
	}

	_, err = fd.Write(bs)
	if err != nil {
		return
	}

	return
}

// Load students info to the file that specified
func Load(fileName string) (stu *StudentSet, err error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	err = json.Unmarshal(bs, &stu)
	if err != nil {
		return
	}

	return
}

func main() {

	var (
		cmd  string
		name string
		file string
		id   int
		line string
		err  error
	)

	f := bufio.NewReader(os.Stdin)
	stus := NewStudentSets()

	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Println(stus)
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			err = stus.Add(id, name)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("add done.\n")
			}
		case "save":
			fmt.Sscan(line, &cmd, &file)
			err = Dump(file, stus)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("save student info to %s done\n", file)
			}
		case "load":
			fmt.Sscan(line, &cmd, &file)
			stus, err = Load(file)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("load student info from %s done\n", file)
			}
		case "remove":
			fmt.Sscan(line, &cmd, &name)
			stus.Remove(name)
			fmt.Printf("remove student %s done\n", name)
		case "clear":
			fmt.Sscan(line, &cmd)
			stus.Clear()
			fmt.Printf("clear all students info %s done\n")
		case "exit", "quit", "q":
			os.Exit(0)
		default:
			fmt.Println("Usage: list|add|save|load|remove|clear|exit")
		}
	}

}
