package main

import (
	"bufio"
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
	var cmd string
	var name string
	var id int
	var line string
	var filename string
	stu_info := make(map[string]Student)
	f := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			for _, v := range stu_info {
				fmt.Println(v.Name, v.Id)
			}

		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			if _, ok := stu_info[name]; ok {
				fmt.Printf("您输入的名字%s,已经存在，请重新输入\n", name)
				return
			} else {
				stu_info[name] = Student{
					Id:   id,
					Name: name}
			}
			fmt.Printf("add done.\n")

		case "save":
			fmt.Sscan(line, &cmd, &filename)
			buf, err := json.Marshal(stu_info)
			if err != nil {
				log.Fatal("marshal err :%s", err)
			}

			f, err := os.Create(filename)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprint(f, string(buf))
			fmt.Printf("save done. finename is %s \n", filename)

		case "load":
			fmt.Sscan(line, &cmd, &filename)
			f, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
				return
			}

			j_err := json.Unmarshal(f, &stu_info)
			if j_err != nil {
				log.Fatal("Unmarshal err :%s", j_err)
			}

			fmt.Printf("load done.\n")

		default:
			fmt.Println("输入有误，请输入以下list|add|save|load")

		}
	}
}
