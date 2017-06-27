package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Student struct for student info
type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var students = make(map[string]Student)

func main() {
	var line, cmd, name string
	var id int
	saved := true

	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		line = strings.Replace(line, "\n", "", 1)
		_, err := fmt.Sscan(line, &cmd, &name, &id)
		// EOF for no enough space-separated values, like `list` cmd
		if err != nil && err != io.EOF {
			fmt.Println(" + parse cmd or info err:", err)
			continue
		}
		if cmd == "exit" {
			break
		}
		switch cmd {
		case "add":
			if addStu(name, id) {
				saved = false
			}
		case "list":
			listStu()
		case "load":
			if loadStu(name, saved) {
				saved = true
			}
		case "save":
			if saveStu(name) {
				saved = true
			}
		case "":
			continue
		default:
			usage()
		}
		line, cmd, name = "", "", ""
	}
}

func saveStu(name string) (rt bool) {
	var err error
	var fd *os.File

	// file existed and override
	if checkFileExist(name) {
		if !checkYes(fmt.Sprintf("override file %s", name)) {
			fmt.Println(" + cancel save to", name)
			return
		}
		if fd, err = os.OpenFile(name, os.O_RDWR|os.O_TRUNC, 0644); err != nil {
			fmt.Printf(" + open file error of %s\n", name)
			return
		}
	}
	// save to new file
	if fd, err = os.Create(name); err != nil {
		fmt.Printf(" + open new file error of %s\n", name)
		return
	}
	if buf, err := json.Marshal(students); err != nil {
		fmt.Println(" + marshal stu info error")
		return
	} else if _, err := fd.Write(buf); err == nil {
		fmt.Println(" + save success")
		return true
	} else {
		fmt.Println(" + save error")
		return
	}
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

// check yes or no for the give question
func checkYes(f string) bool {
	var s string
	for {
		fmt.Printf(" + %s, y or n? :", f)
		fmt.Scanf("%s", &s)
		if string(s[0]) == "y" || string(s[0]) == "Y" {
			return true
		} else if string(s[0]) == "n" || string(s[0]) == "N" {
			return false
		}
	}
}

func loadStu(name string, saved bool) (rt bool) {
	var err error
	var buf []byte
	if !checkFileExist(name) {
		fmt.Printf(" + file : %s not existed!\n", name)
		return
	}
	if !saved && !checkYes("clear up stu info in mem for load") {
		fmt.Println(" + give up load!")
		return
	}
	if buf, err = ioutil.ReadFile(name); err != nil {
		fmt.Printf("read from file error of %s\n", name)
		return
	}
	if err = json.Unmarshal(buf, &students); err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(" + load success")
	return true
}

func addStu(name string, id int) (rt bool) {
	if _, ok := students[name]; ok {
		fmt.Printf(" + duplicated name: %s\n", name)
		return
	}
	students[name] = Student{ID: id, Name: name}
	if _, ok := students[name]; ok {
		fmt.Println(" + add success!")
		return true
	}
	return
}

func listStu() {
	if len(students) == 0 {
		fmt.Println(" + no student info here")
		return
	}
	fmt.Println(" + Id\tName:")
	for _, val := range students {
		fmt.Printf(" + %d\t%s\n", val.ID, val.Name)
	}
}

func usage() {
	fmt.Println(" + cli usage:")
	fmt.Println(" + add name id -- add student info")
	fmt.Println(" + list \t-- list student info")
	fmt.Println(" + load file \t-- load student from file")
	fmt.Println(" + save file \t-- save student info file")
	fmt.Println(" + exit \t-- exit the cli")
}
