package grade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/51reboot/golang-01-homework/lesson9/jungle85gopy/stuServer/class"
)

// Grade map from class name to it's data
type Grade struct {
	// allClasses map from class name to data
	allClasses map[string]*class.Class
}

func (g *Grade) searchClass(name string) bool {
	if len(name) == 0 {
		return false
	}
	for curName := range g.allClasses {
		if name == curName {
			return true
		}
	}
	return false
}

// init to init the allClasses map
func (g *Grade) init() {
	g.allClasses = make(map[string]*class.Class)
}

// PrintGrade print the grade info of all class
func (g *Grade) PrintGrade() {
	log.Printf("grade info:\n")
	for cName, cMap := range g.allClasses {
		log.Printf("%s %v\n", cName, cMap)
	}
}

// Create to creat a new Class
func (g *Grade) Create(_ string, args []string) ([]byte, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("args not enougth")
	}
	name := args[0]
	if existed := g.searchClass(name); existed {
		return nil, fmt.Errorf("class name: %s existed", name)
	}
	if len(g.allClasses) == 0 {
		g.init()
	}
	g.allClasses[name] = new(class.Class)
	g.allClasses[name].SetName(name)
	g.PrintGrade()
	return class.OK, nil
}

// Change changes the current class
func (g *Grade) Change(cur string, args []string) ([]byte, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("args not enougth")
	}
	name := args[0]
	if existed := g.searchClass(name); !existed {
		return g.Create(cur, args)
	}
	return class.OK, nil
}

// Show show all classes
func (g *Grade) Show(_ string, _ []string) ([]byte, error) {
	if len(g.allClasses) == 0 {
		return []byte("no class here\n"), nil
	}
	retStr := "class name info:\n"
	for cName := range g.allClasses {
		retStr += fmt.Sprintf("\t%s\n", cName)
	}
	return []byte(retStr), nil
}

// MarshalJSON my own marshal. cls must be a struct copy. not pointer
func (g Grade) MarshalJSON() ([]byte, error) {
	// tmpStruct := make(map[string]interface{})
	// tmpStruct["AllClasses"] = g.allClasses
	// return json.Marshal(tmpStruct)
	return json.Marshal(struct {
		AllClasses map[string]*class.Class `json:"classes"`
	}{
		AllClasses: g.allClasses,
	})
}

// UnmarshalJSON unmarshal. cls must be a pointer of Student struct
func (g *Grade) UnmarshalJSON(data []byte) error {
	var tmp struct {
		AllClasses map[string]*class.Class `json:"classes"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		log.Println("err of UN,", err)
		return err
	}
	g.allClasses = tmp.AllClasses
	return nil
}

// ------- implement all Method for Class -------

// Add add student info to current class
func (g *Grade) Add(cur string, args []string) ([]byte, error) {
	if cur == "" {
		return nil, fmt.Errorf("current class is blank")
	}
	return g.allClasses[cur].Add(args)
}

// List list student info of current class
func (g *Grade) List(cur string, args []string) ([]byte, error) {
	if cur == "" {
		return nil, fmt.Errorf("current class is blank")
	}
	return g.allClasses[cur].List(args)
}

// Update update student info of current class
func (g *Grade) Update(cur string, args []string) ([]byte, error) {
	if cur == "" {
		return nil, fmt.Errorf("current class is blank")
	}
	return g.allClasses[cur].Update(args)
}

// Delete delete student from current class
func (g *Grade) Delete(cur string, args []string) ([]byte, error) {
	if cur == "" {
		return nil, fmt.Errorf("current class is blank")
	}
	return g.allClasses[cur].Delete(args)
}

// Save save all class info
func (g *Grade) Save(_ string, args []string) ([]byte, error) {
	var err error
	var fd *os.File
	if len(args) < 1 {
		return nil, fmt.Errorf("no save file")
	}
	name := args[0]
	if fd, err = os.Create(name); err != nil {
		return nil, fmt.Errorf("open new file error of %s", name)
	}
	defer fd.Close()

	// saving
	buf, err := json.Marshal(g)
	if err != nil {
		return nil, fmt.Errorf("marshal stu info error")
	}
	if _, err := fd.Write(buf); err != nil {
		return nil, fmt.Errorf("saving error")
	}
	return class.OK, nil
}

// Load load grade info from the given file
func (g *Grade) Load(_ string, args []string) ([]byte, error) {
	var err error
	var buf []byte

	if len(args) < 1 {
		return nil, fmt.Errorf("args not enougth")
	}
	name := args[0]
	if !checkFileExist(name) {
		return nil, fmt.Errorf("%s not existed", name)
	}
	if buf, err = ioutil.ReadFile(name); err != nil {
		return nil, fmt.Errorf("read from file error")
	}
	if err := json.Unmarshal(buf, g); err != nil {
		return nil, fmt.Errorf("unmarshal stu error")
	}
	log.Print("load success")
	return class.OK, nil
}

// check file f exist or not
func checkFileExist(f string) bool {
	if _, err := os.Stat(f); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}
