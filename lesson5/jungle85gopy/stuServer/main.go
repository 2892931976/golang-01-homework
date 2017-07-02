package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// Student struct for student info
type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// original stuMap map from string to Student
// new stuMap map from string to *Student for error: cannot assign
// refer: https://stackoverflow.com/questions/32751537/
//   why-do-i-get-a-cannot-assign-error-when-setting-value-to-a-struct-as-a-value-i
var stuMap = make(map[string]*Student)
var saved bool

func main() {
	actionMap := map[string]func([]string) error{
		"add":    addInfo,
		"list":   listInfo,
		"save":   saveInfo,
		"load":   loadInfo,
		"update": updateInfo,
		"delete": deleteInfo,
		"exit":   exitInfo,
	}
	saved = true
	f := bufio.NewReader(os.Stdin)
	for {
		cmd, args := parseCmd(f)
		if cmd == "" {
			continue
		}
		actionFunc, ok := actionMap[cmd]
		if !ok {
			printInfo(fmt.Sprintf("bad cmd : %s", cmd))
			usage()
			continue
		}
		err := actionFunc(args)
		if err != nil {
			printInfo(fmt.Sprintf("execute action [%s] error: %s", cmd, err))
			continue
		}
	}
	return
}

func parseCmd(f *bufio.Reader) (cmd string, args []string) {
	printInfo(fmt.Sprintf("saved value: %v", saved))
	fmt.Print("> ")
	line, _ := f.ReadString('\n')
	line = strings.TrimSpace(line)
	argsAll := strings.Fields(line)
	if len(argsAll) == 0 {
		return
	} else if len(argsAll) == 1 {
		return argsAll[0], nil
	}
	return argsAll[0], argsAll[1:]
}

func addInfo(args []string) error {
	printInfo("call add, args: " + strings.Join(args, ", "))
	if len(args) <= 1 {
		return fmt.Errorf("args not enougth")
	}
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("error params for id")
	}
	if _, ok := stuMap[name]; ok {
		return fmt.Errorf("duplicated name: " + name)
	}
	stuMap[name] = &Student{ID: id, Name: name}
	if _, ok := stuMap[name]; !ok {
		return fmt.Errorf("add %s failed!" + name)
	}
	saved = false
	return nil
}

func listInfo(args []string) error {
	if len(stuMap) == 0 {
		return fmt.Errorf("no student info here")
	}
	fmt.Println(" + Id\tName:")
	for _, val := range stuMap {
		fmt.Printf(" + %d\t%s\n", val.ID, val.Name)
	}
	return nil
}

func saveInfo(args []string) error {
	var err error
	var fd *os.File
	if len(args) < 1 {
		return fmt.Errorf("args not enougth")
	}
	name := args[0]
	// file existed and override
	if checkFileExist(name) {
		if !checkYes(fmt.Sprintf("override file %s", name)) {
			return fmt.Errorf("cancel save to file " + name)
		}
		if fd, err = os.OpenFile(name, os.O_RDWR|os.O_TRUNC, 0644); err != nil {
			return fmt.Errorf("open file error of " + name)
		}
		defer fd.Close()
	}
	// save to new file
	if fd, err = os.Create(name); err != nil {
		return fmt.Errorf("open new file error of " + name)
	}
	defer fd.Close()

	// saving
	buf, err := json.Marshal(stuMap)
	if err != nil {
		return fmt.Errorf("marshal stu info error")
	}
	if _, err := fd.Write(buf); err != nil {
		return fmt.Errorf("saving error")
	}
	saved = true
	return nil
}

func loadInfo(args []string) error {
	var err error
	var buf []byte

	if len(args) < 1 {
		return fmt.Errorf("args not enougth")
	}
	name := args[0]
	if !checkFileExist(name) {
		return fmt.Errorf(name + " not existed")
	}
	if !saved && !checkYes("clear up stu info in mem for load") {
		return fmt.Errorf("give up load")
	}
	if buf, err = ioutil.ReadFile(name); err != nil {
		return fmt.Errorf("read from file error")
	}
	if err := json.Unmarshal(buf, &stuMap); err != nil {
		return fmt.Errorf("unmarshal stu error")
	}
	printInfo("load success")
	saved = true
	return nil
}

func updateInfo(args []string) error {
	printInfo("call update, args: " + strings.Join(args, ", "))
	if len(args) <= 1 {
		return fmt.Errorf("args not enougth")
	}
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("error params for id")
	}
	if _, ok := stuMap[name]; !ok {
		return fmt.Errorf("%s not exist", name)
	}
	stuMap[name].ID = id
	if stuMap[name].ID != id {
		return fmt.Errorf("update failed for " + name)
	}
	saved = false
	return nil
}

func deleteInfo(args []string) error {
	printInfo("call delete, args: " + strings.Join(args, ", "))
	if len(args) <= 1 {
		return fmt.Errorf("args need to be NAME ID")
	}
	name := args[0]
	stu, ok := stuMap[name]
	if !ok {
		return fmt.Errorf("%s not exist", name)
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("error params for id")
	}
	if stu.ID != id || stu.Name != name {
		return fmt.Errorf("given info not matched")
	}
	delete(stuMap, name)
	if _, ok := stuMap[name]; ok {
		return fmt.Errorf("delete failed for " + name)
	}
	saved = false
	return nil
}

func exitInfo(args []string) error {
	if saved || checkYes("exit with saving stu info") {
		os.Exit(0)
	}
	return nil
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
		printInfo(fmt.Sprintf("%s, y or n?", f))
		fmt.Scanf("%s", &s)
		if string(s[0]) == "y" || string(s[0]) == "Y" {
			return true
		} else if string(s[0]) == "n" || string(s[0]) == "N" {
			return false
		}
	}
}

func printInfo(outs string) {
	fmt.Printf("  -- [%s] %s\n", time.Now().String(), outs)
}

func usage() {
	printInfo("cli usage:")
	fmt.Println("  + add name id -- add student info")
	fmt.Println("  + list \t-- list student info")
	fmt.Println("  + load file \t-- load student from file")
	fmt.Println("  + save file \t-- save student info file")
	fmt.Println("  + exit \t-- exit the cli")
}
