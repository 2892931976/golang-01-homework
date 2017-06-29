package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	var f_name, cmd string
	var s_struct Student
	s_info := make(map[string]Student)

	for {
		fmt.Print("> ")
		fmt.Scan(&cmd) //交互式输入 命令
		if cmd == "stop" {
			break
		}

		switch cmd {
		case "list":
			for k, v := range s_info {
				fmt.Println(k, v)
			}
		case "add":
			fmt.Print(">> ")
			fmt.Scan(&s_struct.Name, &s_struct.Id) //交互式输入 命令
			_, ok := s_info[s_struct.Name]         //检验用户name是否存在
			if ok {
				fmt.Printf("%v用户信息已经存在! ", s_struct.Name)
			} else {
				s_info[s_struct.Name] = s_struct
				//var s_info map[string]Student
			}

		case "save":
			buf, err := json.Marshal(s_info) //执行序列化
			fmt.Print("add save to  file name: ")
			fmt.Scan(&f_name)
			f, err := os.Create(f_name)
			if err != nil {
				log.Fatal(err)
			}
			f.WriteString(string(buf))
			f.Close()
		case "load":
			fmt.Print("load  from file name: ")
			fmt.Scan(&f_name)
			buf, err := ioutil.ReadFile(f_name)
			if err != nil {
				fmt.Println(err)
				return
			}
			str := fmt.Sprintln(string(buf))
			erro := json.Unmarshal([]byte(str), &s_info)
			if erro != nil {
				log.Fatalf("unmarshal error:%s", err)
			}
			fmt.Println(s_info)
		default:
			fmt.Println("指令不正确,请输入add | list | save | load |stop")

		}
	}
}
