package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"errors"
	"encoding/json"
	"io/ioutil"
)

type Student struct {
	Id   int
	Name string
}

type ClassRoom struct {
	Class_Name string
	Students []Student
}

//var s_info = make(map[string]Student)
var classrooms = make(map[string]ClassRoom)

var classroom  ClassRoom

//add id name
func add(args []string) error {
	name := args[0]
	id, err := strconv.Atoi(args[1])
	fmt.Println(name,id)
	if err  != nil {
		log.Fatal(err)
	}
	error := classroom.Add(name,id)
	if  error != nil {
		log.Fatal(error)
	}
	return nil
}

func (cr *ClassRoom)Add(name string ,id int) error  {
	for _, value := range cr.Students {
		if value.Name == name {
			err := fmt.Sprintf("%v %v exist\n", cr.Class_Name, name)
			return errors.New(err)
		}
	}
	s := Student{
		Id:   id,
		Name: name,
	}

	cr.Students = append(cr.Students, s)
	fmt.Printf("添加成功\n")
	classrooms[cr.Class_Name] = *cr
	return nil

}

//list
func list(args []string) error {
	//for k, v := range s_info {
	//	fmt.Println(k, v)
	//}
	//return nil

	error := classroom.List()
	if  error != nil {
		log.Fatal(error)
	}
	return nil
}
func (cr *ClassRoom)List() error  {
	for _,value := range cr.Students {
		fmt.Printf("班级:%v -> ID:%v -> Name:%v\n",cr.Class_Name,value.Id,value.Name)
	}
	return nil
}

//selroom name
func selroom(args []string) error {
	name := args[0]
	if v, ok := classrooms[name]; ok {
		classroom = v
	} else {
		classroom = ClassRoom{
			Class_Name:    name,
			Students: []Student{},
		}
		classrooms[name] = classroom
	}
	fmt.Printf("USE %v\n", name)
	return nil

}


////save filename
func save(args []string) error {
	buf, err := json.Marshal(classrooms) //执行序列化
	f, err := os.Create(args[0])
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(string(buf))
	return  nil
}


//load  filename
func load(args []string) error {
	buf, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}
	str := fmt.Sprintln(string(buf))
	erro := json.Unmarshal([]byte(str), &classrooms)
	if erro != nil {
		log.Fatalf("unmarshal error:%s", err)
	}
	fmt.Println(classrooms)
	return nil
}

////update  name  id
//func update(args []string) error {
//	id, err := strconv.Atoi(args[1])
//	if err != nil {
//		log.Fatal(err)
//	}
//	s_info[args[0]] = Student{
//		Id:   id,
//		Name: args[0],
//	}
//	return nil
//}
//
////del  name
//func del(args []string) error {
//	delete(s_info, args[0])
//	return nil
//}

func main() {
	//			actionmap[cmd] actionfunc([]args)
	actionmap := map[string]func([]string) error{
		"add":    add,
		"list":   list,
		"selroom":selroom,
		"load":   load,
		"save":   save,
		//"update": update,
		//"del":del,
	}

	f := bufio.NewReader(os.Stdin)

	//var students map[string]Student
	for {
		fmt.Print("> ")
		line, _ := f.ReadString('\n')
		line = strings.TrimSpace(line) // 去除两端的空格和换行
		args := strings.Fields(line)   // 按空格分割字符串得到字符串列表
		if len(args) == 0 {
			continue
		}
		// 获取命令和参数列表
		cmd := args[0]
		args = args[1:]

		// 获取命令函数
		//			   map[string]  func([]string) error
		actionfunc := actionmap[cmd]
		if actionfunc == nil {
			fmt.Println("bad cmd ", cmd)
			continue
		}
		//func([]string) error
		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error:%s\n", cmd, err)
			continue
		}
	}
}
