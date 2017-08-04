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

var students = make(map[string]Student)

type Student struct {
	Id   int
	Name string
}

func add(args []string) error {
	fmt.Println("call add")
	fmt.Println("args", args)

	if len(args) < 2 {
		return errors.New("command format: add username id")
	}

	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("Id should be number")
	}

	_, ok := students[name]
	if ok {
		fmt.Printf("%v is in the list \n", name)
	}
	students[name] = Student{Id: id, Name: name}
	return nil
}

func del(args []string) error {
	fmt.Println("delete student")
	fmt.Println("args", args)
	if len(args) < 1 {
		return fmt.Errorf("Please input student name")
	}
	name := args[0]

	_, ok := students[name]
	if ok {
		delete(students, args[0])
		fmt.Printf("%v is deleted \n", name)
		return nil
	} else {
		fmt.Printf("%v is not in the list \n", name)
		return nil
	}
}

func update(args []string) error {
	fmt.Println("update student info")
	fmt.Println("args", args)
	if len(args) < 2 {
		return fmt.Errorf("Please input student name and new Id")
	}

	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("Id should be number")
	}

	_, ok := students[name]
	if ok {
		students[name] = Student{Id: id, Name: name}
	} else {
		return fmt.Errorf("This student is not in the list")
	}
	fmt.Printf("%v id is up to date \n", name)
	return nil
}

func list(args []string) error {
	if len(students) != 0 {
		for _, v := range students {
			fmt.Println(v.Name, v.Id)
		}
	} else {
		return errors.New("unimplemention")
	}

	return nil
}

func save(args []string) error {
	if len(args) < 1 {
		fmt.Printf("Input text name for saving students info \n")
		return nil
	}

	if len(students) == 0 {
		return errors.New("Empty students list, no need to be saved")
	}

	txt := args[0]
	info, err := json.Marshal(students)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(txt, info, 0644)
	fmt.Printf("Saved \n")
	return nil
}

func load(args []string) error {
	if len(args) < 1 {
		fmt.Printf("Input text name for loading students info \n")
		return nil
	}

	txt := args[0]
	info, err := ioutil.ReadFile(txt)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(info, &students)

	fmt.Printf("Loaded %v \n", txt)
	return nil
}

func quit(args []string) error {
	fmt.Println("Quit")
	os.Exit(0)
	return nil
}

func main() {
	actionmap := map[string]func([]string) error{
		"add":    add,
		"del":    del,
		"update": update,
		"list":   list,
		"save":   save,
		"load":   load,
		"quit":   quit,
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
