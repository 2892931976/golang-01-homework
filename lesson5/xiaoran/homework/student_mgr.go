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

var stu_info = make(map[string]Student)

func list(args []string) error {
	for _, v := range stu_info {
		fmt.Println(v.Name, v.Id)
	}
	return nil
}

func add(args []string) error {
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	if _, ok := stu_info[name]; ok {
		fmt.Printf("您输入的名字%s,已经存在，请重新输入\n", name)
	} else {
		stu_info[name] = Student{
			Id:   id,
			Name: name}
	}
	fmt.Printf("add done.\n")
	return nil
}

func del(args []string) error {
	name := args[0]
	if _, ok := stu_info[name]; !ok {
		fmt.Printf("您输入的名字%s不存在，请重新输入\n", name)
	} else {
		delete(stu_info, name)
	}
	return nil
}

func update(args []string) error {
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	if _, ok := stu_info[name]; !ok {
		fmt.Printf("您输入的名字%s不存在，请重新输入\n", name)
	} else {
		stu_info[name] = Student{
			Id:   id,
			Name: name}
	}
	fmt.Printf("update done.\n")
	return nil
}

func save(args []string) error {
	filename := args[0]
	buf, err := json.Marshal(stu_info)
	if err != nil {
		log.Fatal("marshal err :%s", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprint(f, string(buf))
	fmt.Printf("save done. finename is %s \n", filename)
	return nil
}

func load(args []string) error {
	filename := args[0]
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	j_err := json.Unmarshal(f, &stu_info)
	if j_err != nil {
		log.Fatal("Unmarshal err :%s", j_err)
	}
	fmt.Printf("load done.\n")
	return nil
}

func main() {
	f := bufio.NewReader(os.Stdin)

	actionmap := map[string]func([]string) error{
		"add":    add,
		"list":   list,
		"save":   save,
		"load":   load,
		"delete": del,
		"update": update,
	}

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
			fmt.Println("输入命令有误，请输入以下命令list|add|delete|update|save|load")
			continue
		}

		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error:%s\n", cmd, err)
			continue
		}
	}
}
