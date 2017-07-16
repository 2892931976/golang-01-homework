package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func save() error {
	buf, err := json.Marshal(classrooms)
	if err != nil {
		return err
	}
	fmt.Println(string(buf))
	return nil
}

func load() error {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		err := fmt.Errorf("%s ,%s", "load faile", err)
		return err
	}
	json.Unmarshal(f, &classrooms)
	return nil
}

func main() {
	classrooms = make(map[string]*ClassRoom) //这里如果使用 := 则变成了main的局部变量
	classrooms1 := &ClassRoom{
		students: make(map[string]*Student),
	}
	classrooms1.Add("binggan", 1)
	classrooms1.Update("binggan", 2)
	classrooms1.List()

	classrooms2 := &ClassRoom{
		students: make(map[string]*Student),
	}
	classrooms2.Add("binggan", 3)
	classrooms2.Update("binggan", 4)
	classrooms2.List()

	classrooms["51reboot"] = classrooms1
	classrooms["51golang"] = classrooms2
	if err := save(); err != nil {
		log.Fatal(err)
	}
}
