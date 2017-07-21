package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	//"log"
	"os"
	"strconv"
	"strings"
	"bufio"
)

type Student struct {
	Id   int
	Name string
}

var classrooms map[string]*ClassRoom

type ClassRoom struct {
	Name string
	students map[string]*Student
}

var currentClassRoom *ClassRoom

func (c *ClassRoom) MarshalJSON() ([]byte, error) {

	return json.Marshal(c.students)
}

func (c *ClassRoom) UnMarshalJSON(buf []byte) error {

	return json.Unmarshal(buf, &c.students)
}

func choose(args []string) error {
	name := args[0]
	if classroom, ok := classrooms[name]; ok {

		currentClassRoom = classroom
	} else {

		currentClassRoom = &ClassRoom{

			Name:     name,
			students: make(map[string]*Student),
		}
		classrooms[name] = currentClassRoom

	}
	fmt.Printf("choice classroom: %v", name)
	return nil

}

func (c *ClassRoom) List() error {

	for _, v := range c.students {
		fmt.Println(v.Name, v.Id)

	}
	return nil
}
func list(args []string) error {
	currentClassRoom.List()
	return nil
}

func (c *ClassRoom) Add(Name string, id int) error {
	if _, ok := c.students[Name]; ok {

		fmt.Printf("%v is exist", Name)
	} else {
		s := &Student{
			Id:   id,
			Name: Name,
		}

		c.students[Name] = s
	}
	return nil

}

func add(args []string) error {

	name := args[0]
	id, _ := strconv.Atoi(args[1])
	currentClassRoom.Add(name, id)
	return nil
}

func (c *ClassRoom) Update(Name string, id int) error {
	if stu, ok := c.students[Name]; ok {

		c.students[Name] = &Student{
			Name: Name,
			Id:   id,
		}
		c.students[Name].Id = id
		stu.Id = id
	} else {
		fmt.Printf("%v is not exist", Name)
	}
	return nil
}
func update(args []string) error {

	Name := args[0]
	id, _ := strconv.Atoi(args[1])
	currentClassRoom.Add(Name, id)
	return nil
}

func (c *ClassRoom) Delete(Name string) error {

	if _, ok := c.students[Name]; ok {

		delete(c.students, Name)
	} else {

		fmt.Printf("%v is not exist", Name)
	}
	return nil

}

func del(args []string) error {

	Name := args[0]
	currentClassRoom.Delete(Name)
	return nil
}
func save(args []string) error {

	buf, err := json.Marshal(classrooms)
	filename, err := os.Create(os.Args[1])
	if err != nil {

		return err
	}
	defer filename.Close()
	filename.WriteString(string(buf))

	//fmt.Println(string(buf))
	return nil

}

// func load() error {

// 	if len(os.Args) != 2 {

// 		fmt.Println("example input : load filename")
// 		return nil
// 	}
// 	filename, err := ioutil.ReadFile(os.Args[1])
// 	if err != nil {

// 		log.Fatal(err)
// 	}
// 	err = json.Unmarshal(filename, &classrooms)
// 	if err != nil {

// 		log.Fatal(err)
// 	}
// 	return nil
// }

func main() {
	classrooms = make(map[string]*ClassRoom)
	operationMap := map[string]func([]string) error{
		"list":   list,
		"add":    add,
		"del":    del,
		"update": update,
		"save":   save,
	}
	for {
		input := bufio.NewReader(os.Stdin)
		fmt.Print(">:")
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)
		cmdline := strings.Fields(line)
		if len(cmdline) == 0 {

			continue
		}
		cmd := cmdline[0]
		cmdline = cmdline[1:]
		ff := operationMap[cmd]

		if ff == nil {
			fmt.Println("input error")
			continue
		}

		err := ff(cmdline)

		if err != nil {
			fmt.Printf( "caouo %v fasheng %v",cmd, err)
			continue
		}

	}

	// classroom1 := &ClassRoom{

	// 	students: make(map[string]*Student),
	// }
	// fmt.Println(classroom1)
	// classroom1.Add(1, "xulei")
	// fmt.Println("students of classroom 51reboot")
	// //classroom1.Add(2,"xulei2")
	// classroom1.List()
	// //save(classroom1, "a.txt")
	// classroom2 := &ClassRoom{

	// 	students: make(map[string]*Student),
	// }
	// classroom2.Add(1, "zhangsan")
	// fmt.Println("students of classroom golang")
	// classroom2.List()
	// classrooms["51reboot"] = classroom1
	// classrooms["golang"] = classroom2

}
