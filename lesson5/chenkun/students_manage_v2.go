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

var students = make(map[string]Student)

func add(args []string) error {
	uname := args[0]
	uid, _ := strconv.Atoi(args[1])
	_, ok := students[uname]
	if ok {
		fmt.Printf("%v 已存在！输入无效！\n", uname)
	} else {
		students[uname] = Student{uid, uname} // map添加学生
	}
	fmt.Printf("add done.\n")
	return nil
}

func list(args []string) error {
	for _, v := range students {
		fmt.Println(v.Name, v.Id)
	}
	return nil
}

func save(args []string) error {
	json_obj, _ := json.Marshal(students) // 序列化，map转换成json
	filename := args[0]
	f, err := os.Create(filename) // 创建文件
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(f, string(json_obj)) // 写入json串
	f.Close()
	fmt.Printf("save to %v successful! \n", filename)
	return nil
}

func load(args []string) error {
	filename := args[0]
	filebuf, err := ioutil.ReadFile(filename) // 读文件
	if err != nil {
		log.Fatal(err)
	}
	jsonerr := json.Unmarshal(filebuf, &students) // 反序列化，json转换成map
	if jsonerr != nil {
		log.Fatal(jsonerr)
	}
	fmt.Printf("load %v successful! \n", filename)
	return nil
}

func Delete(args []string) error {
	uname := args[0]
	_, ok := students[uname]
	if ok {
		delete(students, uname)
		fmt.Printf("%v 删除成功！", uname)
	} else {
		fmt.Printf("%v 不存在! 无法删除", uname)
	}
	return nil
}

func update(args []string) error {
	uname := args[0]
	uid, _ := strconv.Atoi(args[1])
	_, ok := students[uname]
	if ok {
		students[uname] = Student{uid, uname}
		fmt.Printf("%v 修改成功！", uname)
	} else {
		fmt.Printf("%v 不存在! 修改失败", uname)
	}
	return nil
}

func main() {
	actionmap := map[string]func([]string) error{
		"add":    add,
		"list":   list,
		"save":   save,
		"load":   load,
		"delete": Delete,
		"update": update,
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
