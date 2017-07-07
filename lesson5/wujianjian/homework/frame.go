package main

import (
	"bufio"
	"encoding/json"
	_ "errors"
	"fmt"
	"io"
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

//添加信息
func add(args []string) error {
	//fmt.Println("args", args)
	name := args[0]
	id, _ := strconv.Atoi(args[1])
	if _, ok := students[name]; ok {
		fmt.Printf("学生:%v 已经存在,请勿重复添加!!!", name)
	} else {
		students[name] = Student{
			Id:   id,
			Name: name,
		}
		fmt.Printf("学生:%v 添加成功!!!\n", name)

	}
	return nil
}

//展示信息
func list(args []string) error {
	fmt.Println("姓名 编号")
	for _, v := range students {
		if 0 <= v.Id && v.Id <= 9 {
			fmt.Printf("%v %v\n", v.Name, "0"+strconv.Itoa(v.Id))
		} else {
			fmt.Printf("%v %v\n", v.Name, v.Id)
		}

	}
	return nil
	//return errors.New("list")
}

//更新信息
func update(args []string) error {
	name := args[0]
	//id := args[1]
	nid, _ := strconv.Atoi(args[2])
	for k, v := range students {
		fmt.Println(v.Id, v.Name)
		if k == name {
			if v.Name == name {
				students[name] = Student{
					Id:   nid,
					Name: name,
				}
				fmt.Printf("学生: %v 的Id更新为：%v 成功!!!\n", name, nid)
			} else {
				fmt.Printf("学生: %v 不存在!!!\n", name)
			}
		}
	}
	return nil
	//return errors.New("update")
}

//删除信息
func del(args []string) error {
	//fmt.Println("args", args[0])
	name := args[0]
	if _, ok := students[name]; ok {
		delete(students, name)
		fmt.Printf("学生: %v 删除成功!!!\n", name)
		return nil
	}
	return nil
	//return errors.New("delete")
}

//保存信息
func save(args []string) error {
	//fmt.Println("args", args[0])
	filename := args[0]
	f, _ := os.Create(filename)
	for _, v := range students {
		buf, err := json.Marshal(v)
		if err != nil {
			log.Fatalf("marshal error:%v", err)
		}
		fmt.Fprintln(f, string(buf))
	}
	return nil
	//return errors.New("save")
}

//加载信息
func load(args []string) error {
	filename := args[0]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	fmt.Println("姓名 编号")
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		//fmt.Print(line)
		var s Student
		err1 := json.Unmarshal([]byte(line), &s)
		if err1 != nil {
			log.Fatalf("unmarshal error:%s", err)
		}
		students[s.Name] = Student{
			Id:   s.Id,
			Name: s.Name,
		}
		fmt.Printf("%s %d\n", s.Name, s.Id)
		//fmt.Println(s)
	}
	return nil
	//return errors.New("load")
}

//帮助
func help(args []string) error {
	fmt.Println(`
+++++++++++++++++++++++++++++++++++++++
+ Usage:                              +
+     1.展示信息:                     +
+       > list                        +
+     2.添加信息:                     +
+       > add name id                 +
+     3.更新信息:                     +
+       > update name id num          +
+     4.删除信息                      +
+       > delete name                 +
+     5.保存信息到文件:               +
+       > save filename               +
+     6.从文件加载信息:               +
+       > load filename               +
+     7.显示帮助                      +
+       > help                        +
+     8.退出:                         +
+       > exit                        +
+++++++++++++++++++++++++++++++++++++++
`)
	return nil
}

//退出
func exit(args []string) error {
	os.Exit(0)
	return nil
}

func main() {
	actionmap := map[string]func([]string) error{
		"add":    add,
		"list":   list,
		"update": update,
		"delete": del,
		"save":   save,
		"load":   load,
		"help":   help,
		"exit":   exit,
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
