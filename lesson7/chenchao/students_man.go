package homework

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var userFile string = "students.txt"
var classroom = make(map[string]*Classroom)
var currentClassroom *Classroom
var currentName string

type Classroom struct {
	students map[string]*Student
}

type Student struct {
	Name string
	Id   int
}

func (s *Classroom) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["students"] = s.students
	return json.Marshal(m)
}

func (s *Classroom) UnmarshalJSON(buf []byte) error {
	cc := json.Unmarshal(buf, s.students)
	return cc
}

func (s *Classroom) add(name string, id int, new_student *Student) error {

	if _, ok := s.students[name]; ok {
		fmt.Println("student already exist")
		return nil
	}

	s.students[name] = new_student
	fmt.Println("add new student success.")

	return nil
}

func (s *Classroom) list() error {
	for _, value := range s.students {
		fmt.Println(value.Name, value.Id)
	}
	return nil
}

func (s *Classroom) update(name string, new_student *Student) error {
	s.students[name] = new_student
	return nil
}

func (s *Classroom) deleteS(name string) error {
	delete(s.students, name)
	return nil
}

func choose(args []string) error {
	name := args[0]
	if class, ok := classroom[name]; ok {
		currentClassroom = class
		delete(classroom, name)
		currentName = name
	}
	return nil
}

func addclass(args []string) error {
	name := args[0]
	if _, ok := classroom[name]; !ok {
		c := make(map[string]*Student)
		c["classname"] = &Student{Name: name, Id: 0}
		if len(classroom) == 0 {

			classroom = make(map[string]*Classroom)
			//classroom[name] = &Classroom{teacher:"", students:c}
			classroom[name] = &Classroom{students: c}
			fmt.Println("add class successful")
		} else {
			s := make(map[string]*Student)
			classroom[name].students = s
			fmt.Println("add new class successful")
		}

	} else {
		fmt.Println("class already exsit")
	}
	return nil
}

func addStudent(args []string) error {
	name := args[0]
	id, _ := strconv.Atoi(args[1])

	new_student := Student{
		Name: name,
		Id:   id,
	}

	currentClassroom.add(name, id, &new_student)
	fmt.Println("add new student success.")

	return nil
}

func save(args []string) error {
	classroom[currentName] = currentClassroom
	dump_info, err := json.Marshal(classroom)
	if err != nil {
		fmt.Println("json dump data error", err)
		return err
	}
	if err := json.Unmarshal(dump_info, &classroom); err != nil {
		fmt.Println(err)
	}
	f, err := os.Create(userFile)
	if err != nil {
		fmt.Println("create file error: ", err)
		return err
	}
	defer f.Close()
	f.Write(dump_info)
	fmt.Println("write students data complete!!!")
	return nil
}

func list(args []string) error { // 参数为all 展示所有班级  参数为class 展示班级内的学生
	if len(args) == 0 {
		fmt.Println("list cmd example: 'list all' or 'list classname")
		return nil
	}
	arg := args[0]
	if arg == "all" {
		for k, value := range classroom {
			fmt.Println(k, "    ", value.students)
		}
		return nil
	} else {
		if class, ok := classroom[arg]; ok {
			fmt.Println(class.list())
		} else {
			fmt.Println("not found class: ", arg)
		}
		return nil
	}

}

func load(args []string) error {
	var f *os.File
	_, err := os.Stat(userFile)
	if err != nil { // 文件不存在 新建一个
		f, err = os.Create(userFile)
		if err != nil { // 创建失败
			fmt.Println("Create new file error")
			return err
		}
		defer f.Close()
		fmt.Println("init user fil...")
		return nil
	}
	user_f, err := ioutil.ReadFile(userFile) // 文件本身存在 直接打开
	if err != nil {
		fmt.Println("Open file error")
		return err
	}
	//classroom = make(map[string]*Classroom)
	if err := json.Unmarshal(user_f, &classroom); err != nil {
		fmt.Println("json load students info error:", err.Error())
		return err
	}
	fmt.Println("load students successful")
	return nil

}

func main() {

	actionmap := map[string]func([]string) error{
		"list":       list,
		"load":       load,
		"select":     choose,
		"addstudent": addStudent,
		"save":       save,
		"addclass":   addclass,
	}

	f := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>>")
		line, _ := f.ReadString('\n')
		line = strings.TrimSpace(line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		cmd := args[0]
		args = args[1:]

		actionfunc := actionmap[cmd]
		if actionfunc == nil {
			fmt.Println("bad func :", cmd)
			continue
		}
		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error: %s \n", cmd, err)
			continue
		}
	}
}
