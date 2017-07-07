/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

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
	ID   int
	NAME string
}

var student_num = make(map[int]Student) //需要给student_num赋初值，不可以写成"var  student_num map[int]Student"这样，
// 不然会报错“panic: assignment to entry in nil map”

func main() {
	actiondict := map[string]func([]string) error{ //定义用户的输入内容
		"add":    add,
		"list":   list,
		"update": update,
		"delete": drop,
		"save":   save,
		"load":   load,
		//"where":where,

		//"limit":limit,
		//"sort":sort,

		//"exit":exit,

	}
	fmt.Println("学生管理系统迷你版！")
	f := bufio.NewReader(os.Stdin) //用它读取用户输入的内容
	for {
		fmt.Print("请输入:>>>")
		line, _ := f.ReadString('\n')   //将读取的内容按照"\n"换行符来切分，注意里面是单引号哟！
		line = strings.Trim(line, "\n") //表示只脱去换行符："\n",你可以自定义脱去字符，等效于line = strings.TrimSpace(line)
		content := strings.Fields(line) //按照空格将得来的字符串做成一个切片。

		if len(content) == 0 { //脱去空格
			continue
		}

		cmd := content[0]   //定义执行命令的参数，如add,upadte,list....等等
		args := content[1:] //定义要执行的具体内容

		action_func := actiondict[cmd] //定义用户执行的函数
		if action_func == nil {        //如果输入有问题，告知用户用法
			fmt.Println("Usage: {add|list|where|load|upadte|delete|}[int][string]")
			continue
		}

		err := action_func(args)
		if err != nil {
			fmt.Println("您输入的字符串有问题,案例：add 01 bingan")
			continue
		}
	}
}

func add(args []string) error {
	if len(args) != 2 {
		fmt.Println("您输入的字符串有问题,案例：add 01 bingan")
		return nil
	}
	id := args[0]
	student_name := args[1]
	student_id, _ := strconv.Atoi(id)

	for _, s := range student_num {
		if s.ID == student_id {
			fmt.Println("您输入的ID已经存在，请重新输入")
			return nil
		}
	}

	student_num[len(student_num)+1] = Student{student_id, student_name}
	fmt.Println("Add access!!!")
	return nil
}

func list(args []string) error {
	if len(student_num) == 0 {
		fmt.Println("数据库为空，请自行添加相关信息！")
	}
	for i := 1; i <= len(student_num); i++ {
		fmt.Printf("学生姓名是：'%s'，该学生ID是:'%d'\n", student_num[i].NAME, student_num[i].ID)
	}
	return nil
}

func update(args []string) error {
	if len(args) != 2 {
		fmt.Println("您输入的字符串有问题,案例：add 01 bingan")
		return nil
	}
	id := args[0]
	student_name := args[1]
	student_id, _ := strconv.Atoi(id)
	for i, j := range student_num {
		if j.ID == student_id {
			student_num[i] = Student{student_id, student_name}
			return nil
		}
	}
	fmt.Println("你愁啥？学生ID压根就不存在！")
	return nil
}

func drop(args []string) error {
	if len(args) != 1 {
		fmt.Println("你愁啥？改学生ID压根就不存在！")
		return nil
	}
	id := args[0]
	student_id, _ := strconv.Atoi(id)
	for i, j := range student_num {
		//fmt.Println(i)
		//fmt.Println()
		//fmt.Println(student_num)
		//fmt.Println(j.ID)
		if j.ID == student_id {
			delete(student_num, i) //删除map中该id所对应的key值。但是该功能需要完善！
			fmt.Println("delete successfully!")
			return nil
		}
		fmt.Println(student_num)
	}
	fmt.Println("你愁啥？学生ID压根就不存在！")
	return nil
}

func save(args []string) error {
	if len(args) == 0 {
		fmt.Println("请输入您想要保存的文件名，例如：save student.txt")
		return nil
	}
	file_name := args[0]
	//fmt.Println(filename)
	//fmt.Println(student_num)
	f, err := json.Marshal(student_num)

	if err != nil {
		fmt.Println("序列化出错啦！")
	}
	ioutil.WriteFile(file_name, f, 0644)
	fmt.Println("写入成功")
	return nil
}

func load(args []string) error {
	if len(args) != 1 {
		fmt.Println("输入错误，请重新输入.")
		return nil
	}
	file_name := args[0]
	s, _ := ioutil.ReadFile(file_name)
	json.Unmarshal(s, &student_num)
	fmt.Println("读取成功！")
	return nil
}

//
//func list()  {
//	fmt.Println(111)
//}
//
//func where() {
//	fmt.Println(111)
//}
//
//
//func sort() {
//	fmt.Println(111)
//}
//
//func limit() {
//	fmt.Println(111)
//
//}
