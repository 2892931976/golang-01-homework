package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

var DBfile = "info"
var studentInfo []string

func add(args []string) error {
	fmt.Println("call add")
	fmt.Println("args", args)
	if len(args) == 2 {
		for _, line := range studentInfo {
			studentId := strings.Fields(line)
			if len(studentId) > 0 && studentId[0] == args[0] {
				return errors.New("The student is exits")
			}
		}

	}
	temp := strings.Join(args, " ")
	studentInfo = append(studentInfo, temp)
	//fmt.Println(studentInfo)
	return nil
}
func del(args []string) error {
	fmt.Println("call del")
	fmt.Println("arg", args)
	if len(args) == 1 {
		for i, line := range studentInfo {
			studentId := strings.Fields(line)
			if len(studentId) > 0 && studentId[0] == args[0] {
				//studentInfo[i]=""
				fmt.Printf("del %s\n", studentInfo[i])
				studentInfo[i] = ""
			}
		}
	}
	return nil
}

func list(args []string) error {
	//return errors.New("unimplemention")
	for i, line := range studentInfo {
		fmt.Println(i, line)
	}
	return nil
}

func load(args []string) error {
	if len(args) == 1 {
		DBfile = args[0]
	}
	fmt.Printf("it will open %s\n", DBfile)
	input, err := ioutil.ReadFile(DBfile)
	if err != nil {
		log.Fatal(err)

	}
	lines := strings.Split(string(input), "\n")
	studentInfo = lines
	return nil
}

func help(args []string) error {
	if len(args) == 0 {
		fmt.Println("add students id , it only add new students ")
		fmt.Println("del students ,it will del student which you choose ")
		fmt.Println("list , it will show all the students")
		fmt.Println("update student id, it only change id of student ")
		fmt.Println("save file, if the file is null,it will save to which it open")
		fmt.Println("load file, if the file is null,it will open info ")
	}
	return nil
}
func save(args []string) error {
	if len(args) == 1 {
		DBfile = args[0]
	}
	var outfile string
	for _, line := range studentInfo {
		if len(line) > 0 {
			outfile = outfile + line + "\n"
		}
	}
	//fmt.Println(outfile)

	err := ioutil.WriteFile(DBfile, []byte(outfile), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil

}

func main() {
	actionmap := map[string]func([]string) error{
		"add":  add,
		"list": list,
		"load": load,
		"help": help,
		"del":  del,
		"save": save,
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
