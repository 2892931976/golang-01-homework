package main

import (
	"fmt"
	"encoding/json"
)

type Student struct {

	 Id int
	Name string
}
var classrooms map[string]*ClassRoom

type  ClassRoom struct {
	    students map[string]*Student


}


func (c *ClassRoom)MarshalJSON() ([]byte, error) {

	 return json.Marshal(c.students)
}
func (c *ClassRoom) List() {

     for k,v := range c.students {
		    fmt.Println(k,v)

	 }
}

func (c *ClassRoom) Add(id int, Name string)  error{
           s := &Student {
            Id: id,
            Name: Name,     
 
           }

	    c.students[Name]=s
	    return nil

}

func (c *ClassRoom) Update(id int, Name string) error {
	 if stu, ok := c.students[Name]; ok {

		 c.students[Name] = &Student{
			 Name: Name,
			 Id: id,

		 }
		 c.students[Name].Id = id
		 stu.Id = id
		 }else {
		 ///
	 }
    return nil
}


func save() error {

	    buf, err := json.Marshal(classrooms)
	   if err != nil {

		    return err
	   }

	fmt.Println(string(buf))
	return nil

}

func main() {
         classrooms := make(map[string]*ClassRoom)


	      classroom1 := &ClassRoom {

		    students: make(map[string]*Student),
	      }
	classroom1.Add(1,"xulei")
	fmt.Println("students of classroom 51reboot")
	//classroom1.Add(2,"xulei2")
	classroom1.List()
	classroom2 := &ClassRoom {

		students: make(map[string]*Student),
	}
	classroom2.Add(1, "zhangsan")
	fmt.Println("students of classroom golang")
	classroom2.List()
	classrooms["51reboot"] = classroom1
	classrooms["golang"] = classroom2

}
