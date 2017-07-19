package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var classrooms map[string]*ClassRoom
var currentClassRoom *ClassRoom

type Student struct {
	Id   int
	Name string
}

func (s *ClassRoom) MarshalJSON() ([]byte, error) {
	//如果有教室,教师,学生区分 则通过如下方式添加
	//		m := make(map[string] interface{})
	//		m["teacher"] = s.teacher
	//		m["students"] = s.students
	//		return json.Marshal(m)
	//
	return json.Marshal(s.students)
}

func (s *ClassRoom) UnmarshalJSON(buf []byte) error {
	return json.Unmarshal(buf, &s.students)
}

type ClassRoom struct {
	students map[string]*Student
}

func (c *ClassRoom) List() {
	for _, stu := range c.students {
		fmt.Println(stu.Name, stu.Id)
	}
}

func (c *ClassRoom) Add(name string, id int) error {
	if _, ok := c.students[name]; ok {
		err := fmt.Errorf("学生%s已经存在,请检查", name)
		return err
	}
	c.students[name] = &Student{
		Id:   id,
		Name: name,
	}
	fmt.Printf("Add %v is ok \n", c.students[name].Name)
	return nil
}

func (c *ClassRoom) Update(name string, id int) error {
	/*另外一种方法
	if stu,ok := c.students[name]; ok{
		stu.Id = id
	*/
	if _, ok := c.students[name]; ok {
		c.students[name].Id = id
		fmt.Printf("update %s is ok \n", c.students[name].Name)
	} else {
		err := fmt.Errorf("学生%s不存在,请检查", name)
		return err
	}
	return nil
}

func choose(args []string) error {
	name := args[0]
	if classrooms, ok := classrooms[name]; ok {
		currentClassRoom = classrooms
	} else {
		err := fmt.Errorf("%s ,%s", classrooms, "不存在")
		return err
	}
	return nil
}

func add(args []string) error {
	name := ""
	id := 0
	currentClassRoom.Add(name, id)
	return nil
}

func save(args []string) error {
	if len(args) == 0 {
		err := fmt.Errorf("%s", "No file specified")
		return err
	}
	file := args[0]
	buf, err := json.Marshal(classrooms)
	if err != nil {
		return err
	}
	ioutil.WriteFile(file, buf, 0400)
	fmt.Println("save ok")
	return nil
}

func load(args []string) error {
	if len(args) == 0 {
		err := fmt.Errorf("%s", "No file specified")
		return err
	}
	file := args[0]
	f, err := ioutil.ReadFile(file)
	if err != nil {
		err := fmt.Errorf("%s ,%s", "load faile", err)
		return err
	}
	json.Unmarshal(f, &classrooms)
	return nil
}

func update(args []string) error {

	return nil
}

func del(args []string) error {

	return nil
}

func list(args []string) error {

	return nil
}

func exit(args []string) error {
	fmt.Println("Bye Bye")
	os.Exit(0)
	return nil
}

func main() {
	actionmap := map[string]func([]string) error{
		"add":    add,
		"list":   list,
		"save":   save,
		"load":   load,
		"update": update,
		"delete": del,
		"choose": choose,
		"exit":   exit,
	}
	f := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, _ := f.ReadString('\n')
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		cmd := args[0]
		args = args[1:]
		action := actionmap[cmd]
		if action == nil {
			fmt.Println("bad cmd:", cmd)
			continue
		}
		err := action(args)
		if err != nil {
			fmt.Printf("execute action %s error : %s \n", cmd, err)
			continue
		}
	}
}
