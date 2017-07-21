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

//定义全局变量
var classrooms = make(map[string]ClassRoom) //初始化一个classroom
var currentClassRoom ClassRoom

//定义结构体
type ClassRoom struct { //定义教室,包含名字和学生
	Name     string
	Students map[string]*Student
}

type Student struct { //定义学生,包含学生id和名字
	Id   int
	Name string
}

func (c *ClassRoom) List() error {
	if c.Name == "" {
		err := fmt.Errorf("Please choose the classroom first")
		return err
	}
	for _, stu := range c.Students {
		fmt.Println(stu.Name, stu.Id)
	}
	return nil
}

func (c *ClassRoom) Add(name string, id int) error {
	if c.Name == "" {
		err := fmt.Errorf("Please choose the classroom first")
		return err
	}
	if _, ok := c.Students[name]; ok {
		err := fmt.Errorf("学生%s已经存在,请检查", name)
		return err
	}
	c.Students[name] = &Student{
		Id:   id,
		Name: name,
	}
	fmt.Printf("Add %v is ok \n", c.Students[name].Name)
	return nil
}

func (c *ClassRoom) Update(name string, id int) error {
	if c.Name == "" {
		err := fmt.Errorf("Please choose the classroom first")
		return err
	}
	if c.Students[name].Name == name {
		c.Students[name].Id = id
		fmt.Printf("update %s is ok \n", c.Students[name].Name)
	} else {
		err := fmt.Errorf("学生%s不存在,请检查", name)
		return err
	}
	return nil
}

func (c *ClassRoom) Del(name string) error {
	if c.Name == "" {
		err := fmt.Errorf("Please choose the classroom first")
		return err
	}
	if _, ok := c.Students[name]; ok {
		delete(c.Students, name)
		fmt.Printf("del %s is ok \n", name)
	} else {
		err := fmt.Errorf("学生%s不存在,请检查", name)
		return err
	}
	return nil

}

//调用调用对应的函数
func add(args []string) error {
	if len(args) != 3 {
		err := fmt.Errorf("%s", "add number of args is error")
		return err
	}
	name := args[1]
	id, err := strconv.Atoi(args[2])
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

func list(args []string) error {
	if len(args) != 1 {
		err := fmt.Errorf("%s", "Example : list")
		return err
	}
	err := currentClassRoom.List()
	if err != nil {
		return err
	}
	return nil
}

func save(args []string) error {
	if len(args) != 2 {
		err := fmt.Errorf("%s", "No file specified")
		return err
	}
	file := args[1]
	//buf, err := currentClassRoom.MarshalJSON()
	buf, err := json.Marshal(classrooms)
	if err != nil {
		return err
	}
	ioutil.WriteFile(file, buf, 0400)
	fmt.Println("save ok")
	return nil
}

func load(args []string) error {
	if len(args) != 2 {
		err := fmt.Errorf("%s", "No file specified")
		return err
	}
	file := args[1]
	f, err := ioutil.ReadFile(file)
	if err != nil {
		err := fmt.Errorf("%s ,%s", "load faile", err)
		return err
	}
	//currentClassRoom.UnmarshalJSON(f)
	classrooms = make(map[string]ClassRoom)
	json.Unmarshal(f, &classrooms)
	fmt.Println("load file is ok")
	return nil
}

func update(args []string) error {
	if len(args) != 3 {
		err := fmt.Errorf("%s", "Example : update jcui 1")
		return err
	}
	name := args[1]
	id, err := strconv.Atoi(args[2])
	if err != nil {
		err := fmt.Errorf("%s ,%s", "id type is not int", err)
		return err
	}
	err = currentClassRoom.Update(name, id)
	if err != nil {
		return err
	}
	return nil
}

func del(args []string) error {
	if len(args) != 2 {
		err := fmt.Errorf("%s", "Example : delete jcui")
		return err
	}
	name := args[1]
	err := currentClassRoom.Del(name)
	if err != nil {
		return err
	}
	return nil
}

func choose(args []string) error {
	if len(args) != 2 {
		err := fmt.Errorf("%s", "Example : choose reboot")
		return err
	}
	name := args[1]
	if v, ok := classrooms[name]; ok {
		currentClassRoom = v
	} else {
		currentClassRoom = ClassRoom{
			Name:     name,
			Students: make(map[string]*Student),
		}
		classrooms[name] = currentClassRoom
	}
	fmt.Printf("ClassRoom: %v\n", name)
	return nil
}

func listroom(args []string) error {
	if len(args) != 1 {
		err := fmt.Errorf("%s", "Example : listroom")
		return err
	}
	for key := range classrooms {
		fmt.Printf("ClassRoom: %v\n", key)
	}
	return nil
}

func exit(args []string) error {
	if len(args) == 1 {
		fmt.Println("Bye Bye")
		os.Exit(0)
	}
	return nil
}

//主函数
func main() {
	actionmap := map[string]func([]string) error{
		"add":      add,
		"list":     list,
		"save":     save,
		"load":     load,
		"update":   update,
		"delete":   del,
		"choose":   choose,
		"exit":     exit,
		"listroom": listroom,
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
