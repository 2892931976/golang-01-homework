package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

//var students = make(map[string]Student)

var classrooms = make(map[string]ClassRoom)
var _classroom ClassRoom

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ClassRoom struct {
	rName    string
	Students []Student
}

func (c *ClassRoom) list() {

	fmt.Printf("classroom: %v\n", c.rName)
	for _, value := range c.Students {
		fmt.Printf("%v %v\n", value.Id, value.Name)
	}
}

func (c *ClassRoom) add(id int, name string) error {
	for _, value := range c.Students {
		if value.Name == name {
			err := fmt.Sprintf("%v %v exist\n", c.rName, name)
			return errors.New(err)
		}
	}
	s := Student{
		Id:   id,
		Name: name,
	}
	c.Students = append(c.Students, s)
	fmt.Printf("Added %v\n", name)
	classrooms[c.rName] = *c
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
		err := fmt.Sprintf("%v %v not exist\n", c.rName, name)
		return errors.New(err)
	}
	c.Students = append(c.Students[:i], c.Students[i+1:]...)
	fmt.Printf("Delete %v\n", name)
	classrooms[c.rName] = *c
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
		fmt.Printf("%v is updated\n", name)
		classrooms[c.rName] = *c
		return nil
	}
	err := fmt.Sprintf("%v %v not exist\n", c.rName, name)
	return errors.New(err)
}

func selectroom(args ...string) error {
	if len(args) != 2 {
		return errors.New("Please select classroom, if classroom doesn't exsit, will create it automatically")
	}
	name := args[1]
	if v, ok := classrooms[name]; ok {
		_classroom = v
	} else {
		_classroom = ClassRoom{
			rName:    name,
			Students: []Student{},
		}
		classrooms[name] = _classroom
	}
	fmt.Printf("ClassRoom: %v\n", name)
	return nil
}

func add(args ...string) error {
	fmt.Println("call add")
	fmt.Println("args", args)

	if len(args) < 3 {
		return errors.New("command format: add username id")
	}

	name := args[1]
	id, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.New("Id should be number")
	}

	//_, ok := students[name]
	//if ok {
	//	fmt.Printf("%v is in the list \n", name)
	//}
	//students[name] = Student{Id: id, Name: name}

	if err := _classroom.add(id, name); err != nil {
		return err
	}
	return nil
}

func del(args ...string) error {
	fmt.Println("delete student")
	fmt.Println("args", args)
	if len(args) < 2 {
		return fmt.Errorf("Please input student name")
	}
	name := args[1]

	//	_, ok := students[name]
	//	if ok {
	//		delete(students, args[0])
	//		fmt.Printf("%v is deleted \n", name)
	//		return nil
	//	} else {
	//		fmt.Printf("%v is not in the list \n", name)
	//		return nil
	//	}

	if err := _classroom.del(name); err != nil {
		return err
	}
	return nil

}

func update(args ...string) error {
	fmt.Println("update student info")
	fmt.Println("args", args)
	if len(args) < 3 {
		return fmt.Errorf("Please input student name and new Id")
	}

	name := args[1]
	id, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.New("Id should be number")
	}

	//_, ok := students[name]
	//	if ok {
	//		students[name] = Student{Id: id, Name: name}
	//	} else {
	//		return fmt.Errorf("This student is not in the list")
	//	}
	//	fmt.Printf("%v id is up to date \n", name)

	if err := _classroom.update(id, name); err != nil {
		return err
	}
	return nil
}

func list(args ...string) error {

	//	if len(students) != 0 {
	//		for _, v := range students {
	//			fmt.Println(v.Name, v.Id)
	//		}
	//	} else {
	//		return errors.New("unimplemention")
	//	}

	_classroom.list()
	return nil
}

func save(args ...string) error {
	if len(args) < 2 {
		fmt.Printf("Input text name for saving students info \n")
		return nil
	}

	if len(classrooms) == 0 {
		return errors.New("Empty students list, no need to be saved")
	}

	txt := args[1]
	info, err := json.Marshal(classrooms)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(txt, info, 0644)
	fmt.Printf("Saved \n")
	return nil
}

func load(args ...string) error {
	if len(args) < 2 {
		fmt.Printf("Input text name for loading students info \n")
		return nil
	}

	txt := args[1]
	info, err := ioutil.ReadFile(txt)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(info, &classrooms)
	if err != nil {
		return err
	}

	fmt.Printf("Loaded %v \n", txt)
	return nil
}

func quit(args ...string) error {
	fmt.Println("Quit")
	os.Exit(0)
	return nil
}

func main() {
	actionmap := map[string]func(...string) error{
		"add":        add,
		"del":        del,
		"update":     update,
		"selectroom": selectroom,
		"list":       list,
		"save":       save,
		"load":       load,
		"quit":       quit,
	}

	f := bufio.NewReader(os.Stdin)

	//var students map[string]Student
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
		if _cmd, ok := actionmap[args[0]]; !ok {
			fmt.Println("Usage: add, del, update, selectroom, list, save, load, quit")
			continue
		} else {
			err := _cmd(args...)
			if err != nil {
				fmt.Println(err)
			}
		}
		//cmd := args[0]
		//args = args[1:]

		// 获取命令函数
		//actionfunc := actionmap[cmd]
		//if actionfunc == nil {
		//	fmt.Println("bad cmd ", cmd)
		//	continue
		//}
		//err := actionfunc(args)
		//if err != nil {
		//	fmt.Printf("execute action %s error:%s\n", cmd, err)
		//	continue
		//}
	}
}
