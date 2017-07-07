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

var s_info = make(map[string]Student)

//add id name
func add(args []string) error {

	_, ok := s_info[args[1]] //检验用户name是否存在
	if ok {
		fmt.Printf("%v用户信息已经存在! ", args[1])
	} else {

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err) //打印错误,退出程序
		}
		s_info[args[1]] = Student{Id: id, Name: args[1]}
		//var s_info map[string]Student
	}
	return nil
}

//list
func list(args []string) error {
	for k, v := range s_info {
		fmt.Println(k, v)
	}
	return nil
}

//save filename
func save(args []string) error {
	buf, err := json.Marshal(s_info) //执行序列化
	f, err := os.Create(args[0])
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(string(buf))
	f.Close()

	return nil
}

//load  filename
func load(args []string) error {
	buf, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}
	str := fmt.Sprintln(string(buf))
	erro := json.Unmarshal([]byte(str), &s_info)
	if erro != nil {
		log.Fatalf("unmarshal error:%s", err)
	}
	fmt.Println(s_info)
	return nil
}

//update  name  id
func update(args []string) error {
	id, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	s_info[args[0]] = Student{
		Id:   id,
		Name: args[0],
	}
	return nil
}

//del  name
func del(args []string) error {
	delete(s_info, args[0])
	return nil
}

func main() {
	//			actionmap[cmd] actionfunc([]args)
	actionmap := map[string]func([]string) error{
		"add":    add,
		"list":   list,
		"load":   load,
		"save":   save,
		"update": update,
		"del":    del,
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
