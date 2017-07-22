package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

type ClassRoom struct {
	Name     string
	students map[string]*Student
}

var classrooms map[string]*ClassRoom

var currentClassRoom *ClassRoom

//1.选择教室
func choose(args []string) error {
	name := args[0]
	if classroom, ok := classrooms[name]; ok {
		currentClassRoom = classroom
	} else {
		currentClassRoom = &ClassRoom{
			Name:     name,
			students: make(map[string]*Student),
		}
		classrooms[name] = currentClassRoom
	}
	fmt.Printf("选择教室:%v\n", name)
	return nil
}

//2.添加信息
func (c *ClassRoom) Add(name string, id int) error {
	if _, ok := c.students[name]; ok {
		fmt.Printf("学生: % s已经存在，请勿重复添加!!!", name)
	} else {
		c.students[name] = &Student{
			Id:   id,
			Name: name,
		}
		fmt.Printf("学生: %v 添加成功!!!\n", name)

	}
	classrooms[name] = c
	return nil
}

func add(args []string) error {
	name := args[0]
	id, _ := strconv.Atoi(args[1])
	currentClassRoom.Add(name, id)
	return nil
}

//3.显示信息
func (c *ClassRoom) List() error {
	for _, v := range c.students {
		fmt.Printf("%v %v\n", v.Name, v.Id)
	}
	return nil
}

func list(args []string) error {
	currentClassRoom.List()
	return nil
}

//4.更新信息
func (c *ClassRoom) Update(name string, id int) error {
	fmt.Println(name, id)
	if stu, ok := c.students[name]; ok {
		stu.Id = id
		fmt.Printf("学生: %v 更新成功!!!\n", name)
	}
	return nil
}

func upd(args []string) error {
	name := args[0]
	id, _ := strconv.Atoi(args[1])
	currentClassRoom.Update(name, id)
	return nil
}

//5.删除信息
func (c *ClassRoom) Delete(name string) error {
	if _, ok := c.students[name]; ok {
		delete(c.students, name)
		fmt.Printf("学生: %v 删除成功!!!\n", name)
	}
	return nil
}

func del(args []string) error {
	name := args[0]
	currentClassRoom.Delete(name)
	return nil
}

//6.保存信息
func save(args []string) error {
	//filename := args[0]
	//f, _ := os.Create(filename)
	fmt.Println(classrooms)

	buf, err := json.Marshal(classrooms)
	if err != nil {
		return err
	}
	fmt.Println(string(buf))

	return nil
}

//7.加载信息

//8.显示帮助
func help(args []string) error {
	fmt.Println(`
+++++++++++++++++++++++++++++++++++++++
+ Usage:                              +
+     1.选择教室:                     +
+       > select classroom            +
+     2.展示信息:                     +
+       > list                        +
+     3.添加信息:                     +
+       > add name id                 +
+     4.更新信息:                     +
+       > update name id num          +
+     5.删除信息                      +
+       > delete name                 +
+     6.保存信息到文件:               +
+       > save filename               +
+     7.从文件加载信息:               +
+       > load filename               +
+     8.显示帮助                      +
+       > help                        +
+     9.退出:                         +
+       > exit                        +
+++++++++++++++++++++++++++++++++++++++
`)
	return nil
}

//9.退出
func exit(args []string) error {
	os.Exit(0)
	return nil
}

func main() {
	classrooms = make(map[string]*ClassRoom)

	actionmap := map[string]func([]string) error{
		"select": choose,
		"add":    add,
		"list":   list,
		"update": upd,
		"delete": del,
		"save":   save,
		"help":   help,
		"exit":   exit,
	}

	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := f.ReadString('\n')
		line = strings.TrimSpace(line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		cmd := args[0]
		args = args[1:]
		actionfunc := actionmap[cmd]
		if actionfunc == nil {
			fmt.Println("bad cmd", cmd)
			continue
		}
		err := actionfunc(args)

		if err != nil {
			fmt.Printf("execute action %s error:%s\n", cmd, err)
			continue
		}

	}

}
