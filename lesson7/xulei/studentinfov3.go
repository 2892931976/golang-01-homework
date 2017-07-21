package main

import (
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

var classrooms map[string]*ClassRoom

type ClassRoom struct {
	students map[string]*Student
}

func (c *ClassRoom) MarshalJSON() ([]byte, error) {

	return json.Marshal(c.students)
}

func (c *ClassRoom) UnMarshalJSON(buf []byte) error {

	return json.Unmarshal(buf, &c.students)
}

func (c *ClassRoom) List() {

	for k, v := range c.students {
		fmt.Println(k, v)

	}
}

func (c *ClassRoom) Add(id int, Name string) error {
	if _, ok := c.students[Name]; ok {

		fmt.Printf("%v is exist", Name)
	}
	s := &Student{
		Id:   id,
		Name: Name,
	}

	c.students[Name] = s
	return nil

}

func (c *ClassRoom) Update(id int, Name string) error {
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

func (c *ClassRoom) Delete(Name string) error {

	if _, ok := c.students[Name]; ok {

		delete(c.students, Name)
	} else {

		fmt.Printf("%v is not exist", Name)
	}
	return nil

}
func save(args ...[]string) error {

	if len(os.Args) != 2 {

		fmt.Println("example input : classroom.save(filename)")
		return nil
	}
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

func load() error {

	if len(os.Args) != 2 {

		fmt.Println("example input : load filename")
		return nil
	}
	filename, err := ioutil.ReadFile(os.Args[1])
	if err != nil {

		log.Fatal(err)
	}
	err = json.Unmarshal(filename, &classrooms)
	if err != nil {

		log.Fatal(err)
	}
	return nil
}

func main() {
	classrooms := make(map[string]*ClassRoom)

	classroom1 := &ClassRoom{

		students: make(map[string]*Student),
	}
	classroom1.Add(1, "xulei")
	fmt.Println("students of classroom 51reboot")
	//classroom1.Add(2,"xulei2")
	classroom1.List()
	classroom1.Save("a.txt")
	classroom2 := &ClassRoom{

		students: make(map[string]*Student),
	}
	classroom2.Add(1, "zhangsan")
	fmt.Println("students of classroom golang")
	classroom2.List()
	classrooms["51reboot"] = classroom1
	classrooms["golang"] = classroom2

}
