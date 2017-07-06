package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name string
	Age  int
}

var m map[string]Student = make(map[string]Student)
var s Student

func add(args []string) error {
	name := args[0]
	age, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("学生年龄不正确")
	}
	if m[name].Age != 0 {
		return errors.New("学生信息已存在")
	}
	s.Name = name
	s.Age = age
	m[name] = s
	fmt.Printf("add %s %d done\n", name, age)
	return nil
}

func del(args []string) error {
	name := args[0]
	if m[name].Age == 0 {
		return errors.New("学生信息不存在")
	}
	delete(m, name)
	return nil
}

func update(args []string) error {
	name := args[0]
	if m[name].Age == 0 {
		return errors.New("学生信息不存在")
	}
	age, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("学生年龄信息不正确")
	}
	s.Name = name
	s.Age = age
	m[name] = s
	return nil
}

func list(args []string) error {
	for _, v := range m {
		fmt.Printf("Name: %s\t Age: %d\n", v.Name, v.Age)
	}
	return nil
}

func save(args []string) error {
	file := args[0]
	if len(file) == 0 {
		return errors.New("请输入文件名称")
	}
	buf, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(string(buf) + "\n")
	f.Close()
	fmt.Println("save successful")
	return nil
}

func load(args []string) error {
	file := args[0]
	if len(file) == 0 {
		return errors.New("请输入文件名称")
	}
	var ls map[string]Student
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		_err := json.Unmarshal([]byte(line), &ls) //反序列化
		if _err != nil {
			log.Fatal(err)
		}
		for k, v := range ls {
			m[k] = v
		}
	}
	f.Close()
	fmt.Println("load successful")
	return nil
}

func main() {
	actionmap := map[string]func([]string) error{
		"add":    add,
		"list":   list,
		"del":    del,
		"update": update,
		"save":   save,
		"load":   load,
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
		if cmd == "stop" {
			break
		}
		actionfunc := actionmap[cmd]
		if actionfunc == nil {
			fmt.Println("bad cmd ", cmd)
			continue
		}
		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error: %s\n", cmd, err)
			continue
		}
	}
}
