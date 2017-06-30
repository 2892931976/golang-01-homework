package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	var cmd, name string
	var age int
	var line string
	var file string
	var m map[string]Student = make(map[string]Student)

	var s Student
	f := bufio.NewReader(os.Stdin) //读取标准输入
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n') //获取整行
		fmt.Sscan(line, &cmd)        //以空格分隔，分别给指针赋值
		if cmd == "stop" {
			break
		}
		switch cmd {
		case "list":
			for k, v := range m {
				fmt.Printf("Name: %-10s %v\n", k, v)
			}

		case "add":
			fmt.Sscan(line, &cmd, &name, &age)
			if m[name].Name != "" {
				fmt.Println("学生已存在")
				continue
			}
			s.Name = name
			s.Age = age
			m[name] = s
			fmt.Printf("%s %s %d done\n", cmd, name, age)
		case "save":
			fmt.Sscan(line, &cmd, &file)
			studentPick(file, m)

		case "load":
			fmt.Sscan(line, &cmd, &file)
			result := studentLoad(file)
			for _, info := range result { //遍历学生信息
				//	fmt.Printf("Name: %-10s Age: %d\n", result[name].Name, result[name].Age)
				fmt.Printf("Name: %-10s Age: %d\n", info.Name, info.Age) //打印到控制台
			}
		}
	}
}

func studentPick(file string, name map[string]Student) {

	buf, err := json.Marshal(name) //序列化，将由学生姓名为key， 学生信息组成的结构体为值的map变量进行序列化
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(string(buf) + "\n")
	f.Close()
}

func studentLoad(file string) map[string]Student { //声明反序列化函数，并返回map 类型
	var s map[string]Student
	students := make(map[string]Student)
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		_err := json.Unmarshal([]byte(line), &s) //反序列化
		if _err != nil {
			log.Fatal(err)
		}
		for k, v := range s {
			students[k] = v //循环遍历map,将单行数据添加到新map中，以学生姓名为key,学生信息组成的结构体为值
		}
	}
	f.Close()
	return students //返回文件中所有学生信息组成的map
}