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

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var message = make(map[string]Student)

func help(args ...string) error {
	if len(args) != 1 {
		return errors.New("请按如下格式输入：help")
	}
	fmt.Println("命令帮助: \n" +
		"\t显示信息:\tlist\n" +
		"\t增加信息:\tadd [Name] [Id]\n" +
		"\t删除信息:\tdel [Name]\n" +
		"\t修改信息:\tupdate [Name] [Id]\n" +
		"\t保存信息:\tsave [filename]\n" +
		"\t加载信息:\tload [filename]\n" +
		"\t帮助信息:\thelp\n" +
		"\t退出程序:\texit")
	return nil
}

func list(args ...string) error {
	if len(args) != 1 {
		return errors.New("请按如下格式输入：list")
	}
	for _, v := range message {
		fmt.Printf("%-10v\t%v\n", v.Name, v.Id)
	}
	return nil
}

func add(args ...string) error {
	if len(args) != 3 {
		return errors.New("请按如下格式输入：add [Name] [Id]")
	}
	name := args[1]
	id := args[2]

	if _, ok := message[name]; ok {
		err := fmt.Sprintf("\t**%v is Exist!**", name)
		return errors.New(err)
	}

	n, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("[Id]: 请输入整数类型")
	}

	message[name] = Student{
		Name: name,
		Id:   n,
	}
	fmt.Println("\t**Add is Ok!**")
	return nil
}

func del(args ...string) error {
	if len(args) != 2 {
		return errors.New("请按如下格式输入：del [Name]")
	}
	name := args[1]
	if _, ok := message[name]; !ok {
		err := fmt.Sprintf("\t**%v is not Exist!**", name)
		return errors.New(err)
	}
	delete(message, name)
	fmt.Println("\t**Del is Ok!**")
	return nil
}

func update(args ...string) error {
	if len(args) != 3 {
		return errors.New("请按如下格式输入：update [Name] [Id]")
	}

	name := args[1]
	id := args[2]
	if _, ok := message[name]; !ok {
		err := fmt.Sprintf("\t**%v is not Exist!**", name)
		return errors.New(err)
	}

	n, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("[Id]: 请输入整数类型")
	}

	tmp := message[name]
	tmp.Id = n
	message[name] = tmp
	fmt.Printf("\t**Change %v is Ok!**\n", name)
	return nil
}

func save(args ...string) error {
	if len(args) != 2 {
		return errors.New("请按如下格式输入：save [filename]")
	}
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	filename := args[1]
	err = ioutil.WriteFile(filename, b, 0600)
	if err != nil {
		return err
	}
	fmt.Println("\t**Save is Ok!**")
	return nil
}

func load(args ...string) error {
	if len(args) != 2 {
		return errors.New("请按如下格式输入：load [filename]")
	}
	filename := args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	message = make(map[string]Student)
	err = json.Unmarshal(content, &message)
	if err != nil {
		return err
	}
	fmt.Println("\t**Load is Ok!**")
	return nil
}

func exit(args ...string) error {
	if len(args) != 1 {
		return errors.New("请按如下格式输入：exit")
	}
	os.Exit(1)
	return nil
}

func main() {
	funcMap := map[string]func(...string) error{
		"list":   list,
		"add":    add,
		"del":    del,
		"update": update,
		"save":   save,
		"load":   load,
		"help":   help,
		"exit":   exit,
	}

	for {
		f := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		line, _ := f.ReadString('\n')
		cmdline := strings.Fields(line)

		if len(cmdline) == 0 {
			continue
		}

		if f, ok := funcMap[cmdline[0]]; !ok {
			funcMap["help"](cmdline...)
			continue
		} else {
			err := f(cmdline...)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
