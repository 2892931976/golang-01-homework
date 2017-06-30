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
	f := bufio.NewReader(os.Stdin) //读取标准输入
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n') //获取整行
		fmt.Sscan(line, &cmd)        //以空格分隔，将值分别赋给指针
		if cmd == "stop" {
			break
		}
		switch cmd {
		case "list":
			s := studentLoad()
			for k, v := range s {
				fmt.Printf("姓名：%-10s\t年龄：%d\n", k, v)
			}
		case "add":
			fmt.Sscan(line, &cmd, &name, &age)
			y := searchStudent(name)
			if y {
				fmt.Println("学生已存在！！")
				continue
			}
			studentPick(name, age)
			fmt.Printf("%s %s %d done\n", cmd, name, age)
		}
	}
}

func searchStudent(s string) bool {
	var y bool = false
	students := studentLoad()
	for name, _ := range students {
		if name == s {
			y = true
			break
		}
	}
	return y
}

func studentPick(name string, age int) {
	s := Student{
		Name: name,
		Age:  age,
	}

	buf, err := json.Marshal(s) //序列化
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("/Users/yhzhao/Documents/workspace/Go/src/lesson4/Students.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(string(buf) + "\n")
	f.Close()
}

func studentLoad() (result map[string]int) {
	var s Student
	students := make(map[string]int)
	f, err := os.Open("/Users/yhzhao/Documents/workspace/Go/src/lesson4/Students.txt")
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
		students[s.Name] = s.Age
	}
	f.Close()
	return students
}
