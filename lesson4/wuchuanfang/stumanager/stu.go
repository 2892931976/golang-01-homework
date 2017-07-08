package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	var cmd string
	var name string
	var filename string
	var id int
	var line string
	f := bufio.NewReader(os.Stdin)

	var students map[int]Student
	students = make(map[int]Student)
	fmt.Println("欢迎使用学生信息管理系统。")
	for {
		fmt.Print("请输入list | add name id | save filename | load filename\n")
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Sscan(line, &cmd, &name, &id)
			//i := len(students)
			//students[i] = Student{112, "wuchf"}
			if len(students) == 0 {
				fmt.Print("信息为空，请先输入 add name id 增加学生信息\n\n")
			} else {
				for i := 1; i <= len(students); i++ {
					fmt.Printf("%s  %d\n", students[i].Name, students[i].Id)
				}
			}
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			i := len(students)
			//fmt.Println("学生姓名：", name, "学生ID：", id)
			if i == 0 && name != "" && id != 0 {
				students[i+1] = Student{*&id, *&name}
				//fmt.Printf("%d %d  %s\n\n", i, students[i+1].Id, students[i+1].Name)
				fmt.Printf("添加学生信息成功\n\n")
			} else {
				//判断id是否存在
				flag := false
				for j := 1; j <= i; j++ {
					if students[j].Id != int(*&id) {
						flag = true
						//fmt.Printf("添加学生信息成功.\n\n")
					} else {
						flag = false
						break
					}
				}
				if flag {
					students[i+1] = Student{*&id, *&name}
					fmt.Printf("添加学生信息成功.\n\n")
				} else {
					fmt.Printf("输入错误，或学生ID已存在，请重新输入.\n\n")
				}
			}
		case "save":
			fmt.Sscan(line, &cmd, &filename)
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
		case "load":
			fmt.Sscan(line, &cmd, &filename)
			s, err := ioutil.ReadFile(filename)
			if filename == "" {
				fmt.Println("加载失败，请输入文件名！")
			} else if err != nil {
				fmt.Println("加载失败！", err)
			} else {
				json.Unmarshal(s, &students)
				fmt.Printf("读取成功.\n\n")
			}
		case "exit":
			fmt.Print("已退出系统。\n")
			os.Exit(0)
		default:
			fmt.Print("输入错误，")
		}
	}
}
