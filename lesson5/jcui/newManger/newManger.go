package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

var stuinfo = make(map[string]Student)

func add(args []string) error {
	if len(args) != 2 {
		err := fmt.Errorf("%s", "add number of args is error")
		return err
	}
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		err := fmt.Errorf("%s ,%s", "id type is not int", err)
		return err
	}
	if _, ok := stuinfo[name]; ok {
		//fmt.Println("name is already exists")
		err := fmt.Errorf("%s", "name is already exists")
		return err
	}
	stuinfo[name] = Student{Id: id, Name: name}
	fmt.Println("add is ok")
	return nil
}

func list(args []string) error {
	for _, v := range stuinfo {
		fmt.Println(v.Name, v.Id)
	}
	return nil
}

func save(args []string) error {
	if len(args) == 0 {
		err := fmt.Errorf("%s", "No file specified")
		return err
	}
	file := args[0]
	w, err := json.Marshal(stuinfo)
	if err != nil {
		err := fmt.Errorf("%s ,%s", "save faile", err)
		return err
	}
	ioutil.WriteFile(file, w, 0400)
	fmt.Println("save ok")

	return nil
}

func load(args []string) error {
	if len(args) == 0 {
		err := fmt.Errorf("%s", "No file specified")
		return err
	}
	file := args[0]
	//if len(stuinfo) != 0 {
	//	err := fmt.Errorf("%s", "memory has data ,Please save the data first !")
	//	return err
	//}
	f, err := ioutil.ReadFile(file)
	if err != nil {
		err := fmt.Errorf("%s ,%s", "load faile", err)
		return err
	}
	json.Unmarshal(f, &stuinfo)
	fmt.Println("load ok")
	return nil
}

func exit(args []string) error {
	//if len(stuinfo) != 0 {
	//	err := fmt.Errorf("%s", "memory has data ,Please save the data and exit !")
	//	return err
	//}
	fmt.Println("Bye Bye")
	os.Exit(0)
	return nil
}

func update(args []string) error {
	if _, ok := stuinfo[args[0]]; ok {
		name := args[0]
		id, err := strconv.Atoi(args[1])
		if err != nil {
			err := fmt.Errorf("%s ,%s", "id type is not int, update faild", err)
			return err
		}
		stuinfo[args[0]] = Student{Id: id, Name: name}
		fmt.Printf("update %s is ok \n", args[0])
		return nil
	}
	err := fmt.Errorf("%s is not exist, update faild", args[0])
	return err
}

func del(args []string) error {
	if _, ok := stuinfo[args[0]]; ok {
		delete(stuinfo, args[0])
		fmt.Printf("del %s is ok \n", args[0])
		return nil
	}
	err := fmt.Errorf("%s is not exist, delet faild", args[0])
	return err

}

func main() {
	actionmap := map[string]func([]string) error{
		"add":    add,
		"list":   list,
		"save":   save,
		"load":   load,
		"exit":   exit,
		"update": update,
		"delete": del,
	}
	f := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, _ := f.ReadString('\n')
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		cmd := args[0]
		args = args[1:]
		action := actionmap[cmd]
		if action == nil {
			fmt.Println("bad cmd:", cmd)
			continue
		}
		err := action(args)
		if err != nil {
			fmt.Printf("execute action %s error : %s \n", cmd, err)
			continue
		}
	}
}
