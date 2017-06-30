package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Student struct {
	Id   int
	Name string
}

/*
   功能描述:
       1. list, 列出所有的学生信息.
       2. add name id，添加一个学生的信息，如果name有重复，报错.
       3. save filename，保存所有的学生信息到filename指定的文件中.
       4. load filename, 从filename指定的文件中加载学生信息.
*/

func main() {
	var cmd string
	var name string
	var id int
	var line string
	var filename string
	f := bufio.NewReader(os.Stdin)

	//所有的信息存放到map中，来避免重复
	students := make(map[string]Student)
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Println("姓名 编号")
			for _, v := range students {
				if 0 <= v.Id && v.Id <= 9 {
					fmt.Printf("%s %s\n", v.Name, "0"+strconv.Itoa(v.Id))
				} else {
					fmt.Printf("%s %s\n", v.Name, v.Id)
				}
			}
			//fmt.Printf("binggan 01\njack 02\n")
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			if _, ok := students[name]; ok {
				fmt.Printf("学生:%s 已经存在,请勿重复添加!!!", name)
			} else {
				students[name] = Student{
					Id:   id,
					Name: name,
				}

				//fmt.Println(students[name])
				fmt.Printf("学生:%s 添加成功!!!\n", name)
			}

		case "save":
			fmt.Sscan(line, &cmd, &filename)
			f, _ := os.Create(filename)

			for _, v := range students {
				buf, err := json.Marshal(v)
				if err != nil {
					log.Fatalf("marshal error:%v", err)
				}
				fmt.Fprintln(f, string(buf))
			}

		case "load":
			fmt.Sscan(line, &cmd, &filename)
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
				fmt.Printf("%s %d\n", s.Name, s.Id)
				//fmt.Println(s)
			}

		case "exit":
			os.Exit(0)
		default:
			fmt.Println(`
+++++++++++++++++++++++++++++++++++++++
+ Usage:                              +
+     1.展示信息:                     +
+       > list                        +
+     2.添加信息:                     +
+       > add name id                 +
+     3.保存信息到文件:               +
+       > save filename               +
+     4.从文件加载信息:               +
+       > load filename               +
+     5.退出:                         +
+       > exit                        +
+++++++++++++++++++++++++++++++++++++++
			`)
		}
	}
}
