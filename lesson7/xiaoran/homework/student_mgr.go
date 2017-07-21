package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var classrooms map[string]*ClassRoom
var currentClassRoom *ClassRoom

type ClassRoom struct {
	students map[string]*Student
}

type Student struct {
	Id   int
	Name string
}

//方法
func (c *ClassRoom) Add(name string, id int) error {
	if _, ok := c.students[name]; ok {
		fmt.Printf("您输入的名字%s,已经存在，请重新输入\n", name)
	} else {
		c.students[name] = &Student{
			Id:   id,
			Name: name,
		}
	}
	fmt.Printf("add done.\n")
	return nil
}

func (c *ClassRoom) List() {
	for _, stu := range c.students {
		fmt.Println(stu.Name, stu.Id)
	}

}

func (c *ClassRoom) Update(name string, id int) error {
	if stu, ok := c.students[name]; !ok {
		fmt.Printf("您输入的名字%s不存在，请重新输入\n", name)
	} else {
		stu.Id = id
	}
	fmt.Printf("update done.\n")
	return nil
}

func (c *ClassRoom) Delete(name string) error {
	if _, ok := c.students[name]; !ok {
		fmt.Printf("您输入的名字%s不存在，请重新输入\n", name)
	} else {
		delete(c.students, name)
	}
	fmt.Printf("delete done.\n")
	return nil
}

func (c *ClassRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.students)
}

func (c *ClassRoom) UnmarshalJSON(buf []byte) error {
	return json.Unmarshal(buf, &c.students)
}

//函数
func choose(args []string) error {
	name := args[0]
	if classroom, ok := classrooms[name]; ok {
		currentClassRoom = classroom
	} else {
		currentClassRoom = &ClassRoom{
			students: make(map[string]*Student),
		}
		classrooms[name] = currentClassRoom
		fmt.Printf("自动创建classroom:%s\n", name)
	}
	return nil
}

func add(args []string) error {
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	currentClassRoom.Add(name, id)
	return nil
}

func list(args []string) error {
	currentClassRoom.List()
	return nil
}

func update(args []string) error {
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	currentClassRoom.Update(name, id)
	return nil
}

func del(args []string) error {
	name := args[0]
	currentClassRoom.Delete(name)
	return nil
}

func save(args []string) error {
	filename := args[0]
	buf, err := json.Marshal(classrooms)
	if err != nil {
		log.Fatal("marshal err :%s", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprint(f, string(buf))
	fmt.Printf("save done. finename is %s \n", filename)
	return nil
}

func load(args []string) error {
	filename := args[0]
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	j_err := json.Unmarshal(f, &classrooms)
	if j_err != nil {
		log.Fatal("Unmarshal err :%s", j_err)
	}
	fmt.Printf("load done.\n")
	return nil
}

func main() {
	f := bufio.NewReader(os.Stdin)

	classrooms = make(map[string]*ClassRoom)
	actionmap := map[string]func([]string) error{
		"add":    add,
		"list":   list,
		"save":   save,
		"load":   load,
		"delete": del,
		"update": update,
		"choose": choose,
	}

	for {
		fmt.Print("> ")
		line, _ := f.ReadString('\n')

		// 去除两端的空格和换行
		line = strings.TrimSpace(line)
		// 按空格分割字符串得到字符串列表
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		// 获取命令和参数列表
		cmd := args[0]
		args = args[1:]

		// 获取命令函数
		actionfunc := actionmap[cmd]
		if actionfunc == nil {
			fmt.Println("输入命令有误，请输入以下命令choose|list|add|delete|update|save|load")
			continue
		}

		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error:%s\n", cmd, err)
			continue
		}
	}
}
