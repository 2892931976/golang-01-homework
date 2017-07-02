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

var students = make(map[string]Student)

type Student struct {
	Id   int
	Name string
}

func Add(args []string) error {
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	_, ok := students[name]
	if ok {
		return errors.New("student already exists.")
	}
	students[name] = Student{
		Id:   id,
		Name: name,
	}
	return nil
}

func List(args []string) error {
	fmt.Println("Id\tName")
	for _, stu_info := range students {
		fmt.Printf("%d\t%s\n", stu_info.Id, stu_info.Name)
	}
	return nil
}

func Exit(args []string) error {
	fmt.Println("Bye bye!!!")
	os.Exit(0)
	return nil
}

func Save(args []string) error {
	filename := args[0]
	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()
	bs, err := json.Marshal(students)
	if err != nil {
		return err
	}
	_, err = fd.WriteString(string(bs))
	if err != nil {
		return err
	}
	fmt.Printf("Save data to %s finish.\n", filename)
	return nil
}

func Load(args []string) error {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	filename := args[0]
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bs, &students)
	if err != nil {
		return err
	}
	fmt.Printf("Load data finish from %s.\n", filename)
	return nil
}

func Delete(args []string) error {
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	for name, info := range students {
		if info.Id == n {
			delete(students, name)
			fmt.Printf("delete student %v finish.\n", name)
			return nil
		}
	}
	fmt.Println("delete failed, student not found.")
	return nil
}

// update name=aaa where id=4
func Update(args []string) error {
	var num string
	var newname string

	for _, v := range args {
		if v == "where" {
			t1 := strings.Split(args[len(args)-1], "=")
			num = t1[1]

			t2 := strings.Split(args[0], "=")
			newname = t2[len(t2)-1]
		}
	}

	n, err := strconv.Atoi(num)
	if err != nil {
		return err
	}

	for name, info := range students {
		if info.Id == n {
			delete(students, name)
			students[name] = Student{
				Id:   n,
				Name: newname,
			}
		}
	}
	return nil
}

func Help(args []string) error {
	str := `
	load db.txt
	save db.txt
	add name num
	delete num
	list
	update name=monkey where id=num
	exit
	`
	fmt.Println(str)
	return nil
}

func main() {
	actionmap := map[string]func([]string) error{
		"add":    Add,
		"delete": Delete,
		"update": Update,
		"list":   List,
		"exit":   Exit,
		"save":   Save,
		"load":   Load,
		"help":   Help,
	}

	f := bufio.NewReader(os.Stdin)

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
