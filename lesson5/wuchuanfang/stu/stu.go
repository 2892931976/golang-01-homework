package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

//var students map[int]Student
var students = make(map[int]Student)

func add(args []string) error {
	if len(args) != 2 {
		fmt.Println("输入错误，请重新输入.")
		return nil
	}
	//fmt.Println("args", args)
	name := args[0]
	id := args[1]
	ID, _ := strconv.Atoi(id)
	i := len(students)

	for _, s := range students {
		if s.Id == ID {
			fmt.Println("输入错误，学生ID已存在，请重新输入.")
			return nil
		}
	}
	students[i+1] = Student{ID, name}
	fmt.Printf("添加学生信息成功.\n\n")
	return nil
}

func list(args []string) error {

	if len(students) == 0 {
		fmt.Print("信息为空，请先输入 add name id 增加学生信息\n\n")
	} else {
		for i := 1; i <= len(students); i++ {
			fmt.Printf("%s  %d\n", students[i].Name, students[i].Id)
		}
	}
	return nil
}

func load(args []string) error {
	if len(args) != 1 {
		fmt.Println("输入错误，请重新输入.")
		return nil
	}
	filename := args[0]
	s, err := ioutil.ReadFile(filename)
	if filename == "" {
		fmt.Println("加载失败，请输入文件名！")
	} else if err != nil {
		fmt.Println("加载失败！", err)
	} else {
		json.Unmarshal(s, &students)
		fmt.Printf("读取成功.\n\n")
	}
	return nil
}

func save(args []string) error {
	if len(args) > 1 {
		fmt.Println("输入错误，请重新输入.")
		return nil
	}
	filename := args[0]
	s, err := json.Marshal(students)
	if err != nil || len(students) == 0 {
		fmt.Println("保存失败，数据不能为空！", err)
	} else if filename == "" {
		filename = "stu.txt"
		ioutil.WriteFile(filename, s, 0644)
		fmt.Printf("保存成功,文件名默认为stu.txt.\n\n")
	} else {
		ioutil.WriteFile(filename, s, 0644)
		fmt.Printf("保存成功.\n\n")
	}
	return nil
}

func update(args []string) error {
	if len(args) != 2 {
		fmt.Println("输入错误，请重新输入.")
		return nil
	}
	//fmt.Println("args", args)
	name := args[0]
	id := args[1]
	ID, _ := strconv.Atoi(id)
	//i := len(students)

	for j, s := range students {
		if s.Id == ID {
			students[j] = Student{ID, name}
			fmt.Printf("修改学生信息成功.\n\n")
			return nil
		}
	}
	fmt.Printf("输入错误，学生ID不存在，请重新输入.\n\n")
	return nil
}

func del(args []string) error {
	if len(args) != 1 {
		fmt.Println("输入错误，请重新输入.")
		return nil
	}
	id := args[0]
	ID, _ := strconv.Atoi(id)
	for i, s := range students {
		if s.Id == ID {
			fmt.Printf("删除学生信息成功.\n\n")
			delete(students, i)
			return nil
		}
	}
	fmt.Printf("输入错误，学生ID不存在，请重新输入.\n\n")
	return nil
}

func exit(args []string) error {
	fmt.Print("已退出系统。\n")
	os.Exit(0)
	return nil
}

func main() {
	actionmap := map[string]func([]string) error{
		"add":    add,
		"list":   list,
		"load":   load,
		"save":   save,
		"update": update,
		"delete": del,
		"exit":   exit,
	}

	fmt.Println("欢迎使用学生信息管理系统。")
	f := bufio.NewReader(os.Stdin)

	//var students map[string]Student
	for {
		fmt.Print("请输入list | add name id | update name id | delete id | save filename | load filename\n")
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
			fmt.Println("输入错误，请重新输入.")
			continue
		}

		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error:%s\n", cmd, err)
			continue
		}
	}
}
