package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	Id   int
	Name string
}

// list, 列出所有的学生信息

// add name id，添加一个学生的信息，如果name有重复，报错

// save filename，保存所有的学生信息到filename指定的文件中

// load filename, 从filename指定的文件中加载学生信息

func main() {
	var cmd string
	var name string
	var id int
	var line string
	f := bufio.NewReader(os.Stdin)

	//所有的信息存放到map中，来避免重复
	//var students map[string]Student
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Printf("binggan 01\njack 02\n")
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			fmt.Printf("add done.\n")
		case "save":
			fmt.Println()
		case "load":
			fmt.Println()
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
