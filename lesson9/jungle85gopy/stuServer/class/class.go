package class

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"
)

// OK for return
var OK = []byte(" OK\n")

// Student struct for student info. id is unique.
type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Class struct for a group of student in the same class
type Class struct {
	allStus   []*Student `json:"allstudents"`
	className string     `json:"classname"`
	flag      sync.Mutex
}

// mothods for class

// search if the given id existed, return idx or -1
func (cls *Class) search(id int) int {
	for i := 0; i < len(cls.allStus); i++ {
		if cls.allStus[i].ID == id {
			return i
		}
	}
	return -1
}

// SetName to set the class name
func (cls *Class) SetName(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("name is blank")
	}
	cls.className = name
	return nil
}

// Add add a new student
func (cls *Class) Add(args []string) ([]byte, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("args not enougth")
	}
	name, id := args[0], args[1]
	if len(name) == 0 {
		return nil, fmt.Errorf("name is blank")
	}
	nid, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("id error")
	}

	cls.flag.Lock()
	defer cls.flag.Unlock()
	if idx := cls.search(nid); idx != -1 {
		return nil, fmt.Errorf("duplicated record: %s in %s", name, cls.className)
	}
	cls.allStus = append(cls.allStus, &Student{ID: nid, Name: name})
	if idx := cls.search(nid); idx == -1 {
		return nil, fmt.Errorf("add %s failed to class %s", name, cls.className)
	}
	return OK, nil
}

// List list all stu info in a class
func (cls *Class) List([]string) ([]byte, error) {
	if len(cls.className) == 0 {
		return nil, fmt.Errorf("no class exist")
	}
	if len(cls.allStus) == 0 {
		return nil, fmt.Errorf("no student in class %s", cls.className)
	}
	retStr := fmt.Sprintf(" [%s] student info:\n\tName\tID:\n", cls.className)
	for _, val := range cls.allStus {
		retStr += fmt.Sprintf("\t%s\t%d\n", val.Name, val.ID)
	}
	return []byte(retStr), nil
}

// Update update the give name of a stu by id
func (cls *Class) Update(args []string) ([]byte, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("args not enougth")
	}
	name, id := args[0], args[1]
	nid, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("id error")
	}

	cls.flag.Lock()
	defer cls.flag.Unlock()
	idx := cls.search(nid)
	if idx == -1 {
		return nil, fmt.Errorf("no record for %s in %s", name, cls.className)
	}
	cls.allStus[idx].Name = name
	return OK, nil
}

// Delete delete a given stu by id
func (cls *Class) Delete(args []string) ([]byte, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("args not enougth")
	}
	name, id := args[0], args[1]
	nid, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("id error")
	}

	cls.flag.Lock()
	defer cls.flag.Unlock()
	idx := cls.search(nid)
	if idx == -1 {
		return nil, fmt.Errorf("no record for id:%d in %s", nid, cls.className)
	}
	if cls.allStus[idx].Name != name {
		return nil, fmt.Errorf("given name not match to name: %s", name)
	}
	cls.allStus = append(cls.allStus[:idx], cls.allStus[idx+1:]...)
	return OK, nil
}

// MarshalJSON my own marshal.
// cls must be a struct copy. not pointer
func (cls Class) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		AllStus   []*Student `json:"allstudents"`
		ClassName string     `json:"classname"`
	}{
		AllStus:   cls.allStus,
		ClassName: cls.className,
	})
}

// UnmarshalJSON un
// cls must be a pointer of Student struct
func (cls *Class) UnmarshalJSON(data []byte) error {
	var tmp struct {
		AllStus   []*Student `json:"allstudents"`
		ClassName string     `json:"classname"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		log.Println("err of UN,", err)
		return err
	}
	cls.allStus = tmp.AllStus
	cls.className = tmp.ClassName
	return nil
}
