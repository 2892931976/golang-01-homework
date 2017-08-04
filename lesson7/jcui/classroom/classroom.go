package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//这个版本有问题,可以忽略,直接看class2.go即可!!!
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
	m := make(map[string]interface{})
	m[s.name] = s.students
	return json.Marshal(m)
}

func (s *ClassRoom) UnmarshalJSON(buf []byte) error {
	return json.Unmarshal(buf, &s.students)
}

type ClassRoom struct {
	name     string
	students map[string]*Student
}

func (c *ClassRoom) List() {
	if c == nil {
		fmt.Println("Please choose the classroom first")
		return
	}
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

func (c *ClassRoom) Del(name string) error {
	if _, ok := c.students[name]; ok {
		delete(c.students, name)
		fmt.Printf("del %s is ok \n", name)
	} else {
		err := fmt.Errorf("学生%s不存在,请检查", name)
		return err
	}
	return nil

}

// 具体调用的函数

func choose(args []string) error {
	if len(args) != 1 {
		err := fmt.Errorf("%s", "Example : choose reboot")
		return err
	}
	name := args[0]
	if classrooms, ok := classrooms[name]; ok {
		currentClassRoom = classrooms
	} else {
		currentClassRoom = &ClassRoom{
			name:     name,
			students: make(map[string]*Student),
		}
	}
	fmt.Println(currentClassRoom.name)
	fmt.Printf("ClassRoom: %v\n", name)
	return nil
}

func add(args []string) error {
	if len(args) != 2 {
		err := fmt.Errorf("%s", "add number of args is error")
		return err
	}
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		err := fmt.Errorf("%s ,%s", "id type is not int", err)
		return err
	}
	err = currentClassRoom.Add(name, id)
	if err != nil {
		return err
	}
	return nil
}

func save(args []string) error {
	if len(args) != 1 {
		err := fmt.Errorf("%s", "No file specified")
		return err
	}
	file := args[0]
	buf, err := currentClassRoom.MarshalJSON() //这里有个问题,保存数据的时候没存入classroom的信息
	//buf, err := json.Marshal(currentClassRoom)
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
	currentClassRoom.UnmarshalJSON(f) //这里同样的问题,load的时候必须先choose classroom
	//json.Unmarshal(f, &classrooms)
	return nil
}

func update(args []string) error {
	if len(args) != 2 {
		err := fmt.Errorf("%s", "Example : update jcui 1")
		return err
	}
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		err := fmt.Errorf("%s ,%s", "id type is not int", err)
		return err
	}
	currentClassRoom.Update(name, id)
	return nil
}

func del(args []string) error {
	if len(args) != 1 {
		err := fmt.Errorf("%s", "Example : delete jcui")
		return err
	}
	name := args[0]
	currentClassRoom.Del(name)
	return nil
}

func list(args []string) error {
	currentClassRoom.List()
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
