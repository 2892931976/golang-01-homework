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
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var message = make(map[string]Student)

//处理异常输入信息
func handle(line string) (int, bool) {
	l := strings.TrimSpace(line)
	s := strings.Split(l, " ")
	if len(s) == 1 {
		switch s[0] {
		case "list", "add", "del", "change", "help", "exit":
			return len(s), true
		}
	}
	if len(s) == 2 {
		switch s[0] {
		case "save", "load":
			return len(s), true
		}
	}
	return len(s), false
}

func help() {
	fmt.Println("命令帮助: \n" +
		"\t显示信息:\tlist\n" +
		"\t增加信息:\tadd\n" +
		"\t删除信息:\tdel\n" +
		"\t修改信息:\tchange\n" +
		"\t保存信息:\tsave [filename]\n" +
		"\t加载信息:\tload [filename]\n" +
		"\t帮助信息:\thelp\n" +
		"\t退出程序:\texit")
}

func list() {
	for _, v := range message {
		fmt.Printf("%-10v\t%v\n", v.Name, v.Id)
	}
}

func add() {
	var name string
	var id string
	fmt.Println("请按如下格式添加：Name Id")
	fmt.Scan(&name, &id)
	if _, ok := message[name]; ok {
		fmt.Printf("\t**%v is Exist!**\n", name)
		return
	}
	n, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Id not Int Type")
		return
	}
	message[name] = Student{
		Name: name,
		Id:   n,
	}
	fmt.Println("\t**Add is Ok!**")
}

func del() {
	var name string
	fmt.Println("请按如下格式添加：Name")
	fmt.Scan(&name)
	if _, ok := message[name]; ok {
		delete(message, name)
		fmt.Printf("\t**Delete %v is Ok!**\n", name)
	} else {
		fmt.Printf("\t**%v is not Exist!**\n", name)
	}
}

func change() {
	var name string
	var id string
	fmt.Println("请按如下格式添加：Name Id")
	fmt.Scan(&name, &id)
	if _, ok := message[name]; ok {
		n, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Id not Int Type")
			return
		}
		tmp := message[name]
		tmp.Id = n
		message[name] = tmp
		fmt.Printf("\t**Change %v is Ok!**\n", name)
	} else {
		fmt.Printf("\t**%v is not Exist!**\n", name)
	}
}

func save(filename string) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0600)
	if err != nil {
		return err
	}
	fmt.Println("\t**Save is Ok!**")
	return nil
}

func load(filename string) error {
	message = make(map[string]Student)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &message)
	if err != nil {
		return err
	}
	fmt.Println("\t**Load is Ok!**")
	return nil
}

func main() {
	for {
		var cmd string
		var line string
		var filename string
		f := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		line, _ = f.ReadString('\n')

		n, b := handle(line)
		if b {
			if n == 1 {
				fmt.Sscan(line, &cmd)
			} else if n == 2 {
				fmt.Sscan(line, &cmd, &filename)
			}
		} else {
			help()
			continue
		}

		switch cmd {
		case "list":
			list()
		case "add":
			add()
		case "del":
			del()
		case "change":
			change()
		case "save":
			err := save(filename)
			if err != nil {
				log.Println(err)
			}
		case "load":
			err := load(filename)
			if err != nil {
				log.Println(err)
			}
		case "help":
			help()
		case "exit":
			return
		}
	}
}
