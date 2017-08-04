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

type Student struct {
	Id   int
	Name string
}

type ClassRoom struct {
	students map[string]*Student
}

var classrooms map[string]*ClassRoom                       //定义classrooms
var classroom_info = &ClassRoom{make(map[string]*Student)} //初始化一个ClassRoom

//listrooms 列出班级
func listrooms(args []string) error {
	for key, _ := range classrooms {
		fmt.Println(key)
	}
	return nil
}

func list(args []string) error {
	classroom_info.List()
	//for _, value := range classroom_info.students {
	//	fmt.Println(value.Name, value.Id)
	//}
	return nil
}

func (c *ClassRoom) List() {
	for _, value := range c.students {
		if value == nil {
			fmt.Println("")
		}
		fmt.Println(value.Name, value.Id)
	}

}

//add name  id
func add(args []string) error {
	if len(args) != 2 {
		fmt.Println("参数有误,使用方法: add name id")
		return nil
	}
	classroom_info.Add(args)
	return nil
}
func (c *ClassRoom) Add(args []string) error { //如果想要修改时,需要传指针 *
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range c.students {
		if value.Name == name {
			fmt.Println("用户已经存在", name)
			continue
		}
	}
	c.students[name] = &Student{id, name}
	return nil
}

//del name
func del(args []string) error {
	if len(args) != 1 {
		fmt.Println("参数有误,使用方法: del name")
		return nil
	}
	classroom_info.Del(args)
	return nil
}
func (c *ClassRoom) Del(args []string) error { //如果想要修改时,需要传指针 *
	name := args[0]
	_, ok := c.students[name]
	if ok { //存在
		delete(c.students, name)
	} else { //不存在
		fmt.Println("用户名不存在", name)
	}
	return nil
}

//update name id
func update(args []string) error {
	if len(args) != 2 {
		fmt.Println("参数有误,使用方法: update name id")
		return nil
	}
	classroom_info.Update(args)
	return nil
}
func (c *ClassRoom) Update(args []string) error {
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	stu, ok := c.students[name] //前提students map[string]*Student,设置为指针方式
	if ok {                     //存在
		stu.Id = id
	} else { //不存在
		fmt.Println("用户名不存在,不能进行更新", name)
	}
	return nil
}

//selroom name
func selroom(args []string) error {
	if len(args) != 1 {
		fmt.Println("参数有误,使用方法: selroom classname ")
		return nil
	}
	name := args[0]
	if v, ok := classrooms[name]; ok {
		classroom_info = v
	} else {
		classrooms[name] = classroom_info
	}
	fmt.Printf("USE %v\n", name)
	return nil
}

func (c *ClassRoom) MarshalJSON() ([]byte, error) { //自定义MarshalJSON接口,序列化
	return json.Marshal(c.students)
}

func (c *ClassRoom) UnmarshalJSON(buf []byte) error { //自定义MarshalJSON接口,反序列化
	return json.Unmarshal(buf, &c.students)
}

//save info.txt
func save(args []string) error {
	if len(args) != 1 {
		fmt.Println("参数有误,使用方法: save filename ")
		return nil
	}
	buf, err := json.Marshal(classrooms) //执行序列化

	f, err := os.Create(args[0])
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(string(buf))
	return nil
}

//load info.txt
func load(args []string) error {
	if len(args) != 1 {
		fmt.Println("参数有误,使用方法: load filename ")
		return nil
	}
	f, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(f, &classrooms) //load的话必须传递 classrooms的地址
	if err != nil {
		log.Fatalf("反序列化失败", err)
	}
	return nil
}

func main() {
	classrooms = make(map[string]*ClassRoom) //初始化classrooms

	actionmap := map[string]func([]string) error{
		"add":       add,
		"list":      list,
		"selroom":   selroom,
		"load":      load,
		"save":      save,
		"update":    update,
		"del":       del,
		"listrooms": listrooms,
	}

	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := f.ReadString('\n')
		line = strings.TrimSpace(line) // 去除两端的空格和换行
		args := strings.Fields(line)   // 按空格分割字符串得到字符串列表
		if len(args) == 0 {
			continue
		}
		cmd := args[0]
		args = args[1:]

		actionfunc := actionmap[cmd]
		if actionfunc == nil {
			fmt.Println("命令不存在", cmd, "\n使用一下命令:\nadd 添加用户,\nlist 展示用户,\nselroom 选择班级,\nupdate 更新信息,\nload 加载,\nsave 保存,\ndel 删除,\nlistrooms 列出班级\n")
			continue
		}
		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error:%s\n", cmd, err)
			continue
		}
	}

}
