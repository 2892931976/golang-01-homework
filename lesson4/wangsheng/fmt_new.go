package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	//"log"
)

type Student struct {
	Id   int
	Name string
}

var info = make(map[string]Student)

func main() {
	var username string
	var password string
	var strpass string
	var line string
	var cmd string
	var filename string
	var infofile string
	filename = "output.txt"
	infofile = "a.txt"
	f := bufio.NewReader(os.Stdin)
	strpass = "abc123"
	fmt.Printf("欢迎进入学生管理系统\n")
	fmt.Printf("请输入你的用户名:")
	fmt.Scan(&username)
	//fmt.Printf("name is %s",name)
	fmt.Printf("请输入你的密码: 密码是abc123")
	fmt.Scan(&password)

	if password != strpass {
		fmt.Printf("密码不正确请重新输入")
		fmt.Scan(&password)

	}
	if password == strpass {

		for {
			fmt.Print("> ")
			line, _ = f.ReadString('\n')
			fmt.Sscan(line, &cmd)
			switch cmd {
			case "list":
				list()
			case "add":
				add()
			case "save":
				save(filename)
			case "load":
				err := load(infofile)
				if err != nil {
					fmt.Println("+++++++++++++")
				}
			}
		}

	}

}

func add() {
	var name string
	var id string
	//fmt.Println("++++++++++++++++")
	fmt.Scan(&name, &id)
	//fmt.Println("---------------")
	n, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Id not Int Type")
		return
	}
	fmt.Println("------------")
	info[name] = Student{
		Name: name,
		Id:   n,
	}
	fmt.Println("\t--Add is ok--")

}
func list() {
	for _, v := range info {
		fmt.Printf("%-10v\t%v\n", v.Name, v.Id)
	}

}

func save(filename string) error {
	b, err := json.Marshal(info)
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
	fmt.Println(filename)
	info = make(map[string]Student)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &info)
	if err != nil {
		return err
	}
	fmt.Println("\t**Load is Ok!**")
	return nil
}
