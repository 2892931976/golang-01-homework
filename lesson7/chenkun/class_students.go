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

// 类型
type Student struct {
	Name string
	Id   int
}

type ClassRoom struct {
	//teacher  string
	students map[string]*Student
}

// 方法
func (c *ClassRoom) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	//m["teacher"] = c.teacher
	m["students"] = c.students
	return json.Marshal(m)
}

func (c *ClassRoom) UnmarshalJSON(buf []byte) error {
	return json.Unmarshal(buf, &c.students)
}

func (c *ClassRoom) List() {
	for _, stu := range c.students {
		fmt.Println(stu.Name, stu.Id)
	}
}

func (c *ClassRoom) Add(name string, id int) error {
	c.students[name] = &Student{
		Name: name,
		Id:   id,
	}
	return nil
}

func (c *ClassRoom) Update(name string, id int) error {
	if stu, ok := c.students[name]; ok {
		c.students[name] = &Student{
			Name: name,
			Id:   id,
		}

		c.students[name].Id = id
		stu.Id = id
	} else {
		///
	}
	return nil
}

// action func
func save(args []string) error {
	json_obj, err := json.Marshal(classrooms)
	filename := args[0]
	if err != nil {
		return err
	}
	f, err := os.Create(filename) // 创建文件
	if err != nil {
		return err
	}
	fmt.Fprint(f, string(json_obj)) // 写入json串
	f.Close()
	fmt.Printf("save to %v successful! \n", filename)
	return nil
}

func choose(args []string) error {
	name := args[0]
	fmt.Println("选择了班级：", name)
	if classroom, ok := classrooms[name]; ok {
		currentClassRoom = classroom
	} else {
		currentClassRoom = &ClassRoom{make(map[string]*Student)}
		classrooms[name] = currentClassRoom
	}
	return nil
}

func add(args []string) error {
	name := args[0]                //
	id, _ := strconv.Atoi(args[1]) //
	currentClassRoom.Add(name, id)
	fmt.Println(name, " 已添加！")
	return nil
}

func list(args []string) error {
	for _, v := range currentClassRoom.students {
		fmt.Println(v.Name, v.Id)
	}
	return nil
}

func Delete(args []string) error {
	name := args[0]
	_, ok := currentClassRoom.students[name]
	if ok {
		delete(currentClassRoom.students, name)
		fmt.Printf("%v 删除成功！", name)
	} else {
		fmt.Printf("%v 不存在! 无法删除", name)
	}
	return nil
}

func Load(args []string) error {
	filename := args[0]
	filebuf, err := ioutil.ReadFile(filename) // 读文件
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(filebuf, &classrooms) // json转换成map

	fmt.Printf("load %v successful! \n", filename)
	return nil
}

func update(args []string) error {
	name := args[0]                //
	id, _ := strconv.Atoi(args[1]) //
	_, ok := currentClassRoom.students[name]
	if ok {
		currentClassRoom.Update(name, id)
		fmt.Printf("%v 修改成功！", name)
	} else {
		fmt.Printf("%v 不存在! 修改失败", name)
	}
	return nil
}

var actionmap = map[string]func([]string) error{
	"add":    add,
	"list":   list,
	"save":   save,
	"load":   Load, // 未完成，现加载后无法转换成对应的map...
	"delete": Delete,
	"update": update,
	"choose": choose,
}

func main() {
	classrooms = make(map[string]*ClassRoom)
	f := bufio.NewReader(os.Stdin)
	for {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("请先choose班级")
			}
		}()
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
			fmt.Println("bad cmd ", cmd)
			continue
		}
		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error:%s\n", cmd, err)
			continue
		}
	}

}
