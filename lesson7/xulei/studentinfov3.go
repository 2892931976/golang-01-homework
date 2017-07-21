package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

var classrooms map[string]*ClassRoom

type ClassRoom struct {
	students map[string]*Student
}

var currentClassRoom *ClassRoom

func (c *ClassRoom) MarshalJSON() ([]byte, error) {

	return json.Marshal(c.students)
}

func (c *ClassRoom) UnMarshalJSON(buf []byte) error {

	return json.Unmarshal(buf, &c.students)
}

func choose(args ...[]string) error {
	name := args[0]
	if classroom, ok := classrooms[Name]; ok {

		currentClassRoom = classroom
	} else {

		currentClassRoom = &ClassRoom{

			Name:     Name,
			students: make(map[string]*Student),
		}
		classrooms[Name] = currentClassRoom

	}
	fmt.Printf("choice classroom: %v", Name)

}

func (c *ClassRoom) List() error {

	for k, v := range c.students {
		fmt.Println(v.Name, v.Id)

	}
}
func list(args ...[]string) error {
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

func add(args, ...[]string) error {

	Name := args[0]
	id, _ := strconv.Atoi(args[1])
	currentClassRoom.Add(Name, id)
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
func update(args ...[]string) error {

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

func del(args ...[]string) error {

	Name := args[0]
	currentClassRoom.Delete(Name)
	return nil
}
func save(args ...[]string) error {

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
	classrooms := make(map[string]*ClassRoom)
	operationMap := map[string]func(...string) error{
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

		if f == nil {
			fmt.Println("input error")
			continue
		}

		if err := ff(cmdline...); err != nil {
			fmt.Println(err)
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
