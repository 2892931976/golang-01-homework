package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var classrooms = make(map[string]ClassRoom)
var croom ClassRoom

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ClassRoom struct {
	Name     string
	Students []Student
}

func (c *ClassRoom) list() {
	fmt.Printf("ClassRoom: %v\n", c.Name)
	for _, value := range c.Students {
		fmt.Printf("\t%v\t%v\n", value.Id, value.Name)
	}
}

func (c *ClassRoom) add(id int, name string) error {
	for _, value := range c.Students {
		if value.Name == name {
			err := fmt.Sprintf("[ClassRoom: %v] Student: %v is exist\n", c.Name, name)
			return errors.New(err)
		}
	}
	s := Student{
		Id:   id,
		Name: name,
	}
	c.Students = append(c.Students, s)
	fmt.Printf("\t**Add %v is Ok!**\n", name)
	classrooms[c.Name] = *c
	return nil
}

func (c *ClassRoom) del(name string) error {
	i := -1
	for key, value := range c.Students {
		if value.Name == name {
			i = key
			break
		}
	}
	if i == -1 {
		err := fmt.Sprintf("[ClassRoom: %v] Student: %v is not exist\n", c.Name, name)
		return errors.New(err)
	}
	c.Students = append(c.Students[:i], c.Students[i+1:]...)
	fmt.Printf("\t**Del %v is Ok!**\n", name)
	classrooms[c.Name] = *c
	return nil
}

func (c *ClassRoom) update(id int, name string) error {
	flag := false
	for key, value := range c.Students {
		if value.Name == name {
			c.Students[key].Id = id
			flag = true
			break
		}
	}
	if flag {
		fmt.Printf("\t**Update %v is Ok!**\n", name)
		classrooms[c.Name] = *c
		return nil
	}
	err := fmt.Sprintf("[ClassRoom: %v] Student: %v is not exist\n", c.Name, name)
	return errors.New(err)
}

func choice(args ...string) error {
	if len(args) != 2 {
		return errors.New("请按如下格式输入：choice [Name]")
	}
	name := args[1]
	if v, ok := classrooms[name]; ok {
		croom = v
	} else {
		croom = ClassRoom{
			Name:     name,
			Students: []Student{},
		}
		classrooms[name] = croom
	}
	fmt.Printf("ClassRoom: %v\n", name)
	return nil
}

func echo(args ...string) error {
	if len(args) != 1 {
		return errors.New("请按如下格式输入：echo")
	}
	for key := range classrooms {
		fmt.Printf("ClassRoom: %v\n", key)
	}
	return nil
}

func list(args ...string) error {
	if len(args) != 1 {
		return errors.New("请按如下格式输入：list")
	}
	croom.list()
	return nil
}

func add(args ...string) error {
	if len(args) != 3 {
		return errors.New("请按如下格式输入：add [Name] [Id]")
	}
	name := args[1]
	id := args[2]

	n, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("[Id]: 请输入整数类型")
	}

	if err := croom.add(n, name); err != nil {
		return err
	}
	return nil
}

func del(args ...string) error {
	if len(args) != 2 {
		return errors.New("请按如下格式输入：del [Name]")
	}
	name := args[1]
	if err := croom.del(name); err != nil {
		return err
	}
	return nil
}

func update(args ...string) error {
	if len(args) != 3 {
		return errors.New("请按如下格式输入：update [Name] [Id]")
	}

	name := args[1]
	id := args[2]

	n, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("[Id]: 请输入整数类型")
	}

	if err := croom.update(n, name); err != nil {
		return err
	}
	return nil
}

func save(args ...string) error {
	if len(args) != 2 {
		return errors.New("请按如下格式输入：save [filename]")
	}
	b, err := json.Marshal(classrooms)
	if err != nil {
		return err
	}
	filename := args[1]
	err = ioutil.WriteFile(filename, b, 0600)
	if err != nil {
		return err
	}
	fmt.Printf("\t**Save %v is Ok!**\n", filename)
	return nil
}

func load(args ...string) error {
	if len(args) != 2 {
		return errors.New("请按如下格式输入：load [filename]")
	}
	filename := args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	classrooms = make(map[string]ClassRoom)
	err = json.Unmarshal(content, &classrooms)
	if err != nil {
		return err
	}
	name := croom.Name
	croom = classrooms[name]
	fmt.Printf("\t**Load %v is Ok!**\n", filename)
	return nil

}

func help(args ...string) error {
	if len(args) != 1 {
		return errors.New("请按如下格式输入：help")
	}
	fmt.Println("命令帮助提示: \n" +
		"\t选择教室信息:\tchoice [Name]\n" +
		"\t显示教室信息:\techo\n" +
		"\t显示学生信息:\tlist\n" +
		"\t增加学生信息:\tadd [Name] [Id]\n" +
		"\t删除学生信息:\tdel [Name]\n" +
		"\t修改学生信息:\tupdate [Name] [Id]\n" +
		"\t保存信息:\tsave [filename]\n" +
		"\t加载信息:\tload [filename]\n" +
		"\t帮助信息:\thelp\n" +
		"\t退出信息:\texit")
	return nil
}

func exit(args ...string) error {
	if len(args) != 1 {
		return errors.New("请按如下格式输入：exit")
	}
	os.Exit(1)
	return nil
}

func main() {
	funcMap := map[string]func(...string) error{
		"choice": choice,
		"echo":   echo,
		"list":   list,
		"add":    add,
		"del":    del,
		"update": update,
		"save":   save,
		"load":   load,
		"help":   help,
		"exit":   exit,
	}

	for {
		f := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		line, _ := f.ReadString('\n')
		cmdline := strings.Fields(line)

		if len(cmdline) == 0 {
			continue
		}

		if f, ok := funcMap[cmdline[0]]; !ok {
			funcMap["help"](cmdline...)
			continue
		} else {
			err := f(cmdline...)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
