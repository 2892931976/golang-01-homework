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
	var filename string
	var name string
	var id int
	var line string
	finput := bufio.NewReader(os.Stdin)

	students := make(map[string]Student)
	for {
		fmt.Print("> ")
		line, _ = finput.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			for _, v := range students {
				fmt.Println(v.Name, v.Id)
			}
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			// 判断add的学生是否已存在
			_, ok := students[name]
			if ok {
				fmt.Printf("%v 已存在！输入无效！\n", name)
				break
			} else {
				students[name] = Student{id, name} // map添加学生
			}
			fmt.Printf("add done.\n")
		case "save":
			fmt.Sscan(line, &cmd, &filename)
			json_obj, _ := json.Marshal(students) // 序列化，map转换成json
			f, err := os.Create(filename)         // 创建文件
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprint(f, string(json_obj)) // 写入json串
			f.Close()
			fmt.Printf("save to %v successful! \n", filename)
		case "load":
			fmt.Sscan(line, &cmd, &filename)
			filebuf, err := ioutil.ReadFile(filename) // 读文件
			if err != nil {
				fmt.Println(err)
				break
			}
			bf := []byte(filebuf)                    // string 转换成 byte
			jsonerr := json.Unmarshal(bf, &students) // 反序列化，json转换成map
			if jsonerr != nil {
				fmt.Println(jsonerr)
			}
			fmt.Printf("load %v successful! \n", filename)
		default:
			fmt.Println("go ...")
		}
	}
}
