package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/51reboot/golang-01-homework/lesson9/jungle85gopy/stuServer/grade"
)

const routineNum = 8

var usage = `  cli usage:
    > create clsName 	-- add class with name
    > select clsName 	-- change to class or create a new class
    > show		-- show all class
    > add stuName id 	-- add student info
    > list 		-- list student info
    > update stuName id 	-- update student name by id
    > delete stuName id 	-- delete student by name or id
    > load file 	   	-- load from file
    > save file 	   	-- save info file
`

var grd grade.Grade
var actionMap map[string]func(string, []string) ([]byte, error)

func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	var grd grade.Grade
	fmt.Println("start:")

	actionMap = map[string]func(string, []string) ([]byte, error){
		"create": grd.Create,
		"select": grd.Change,
		"show":   grd.Show,

		"add":    grd.Add,
		"list":   grd.List,
		"save":   grd.Save,
		"load":   grd.Load,
		"update": grd.Update,
		"delete": grd.Delete,
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		log.Print("class: new connection.\n")
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	var rd *bufio.Reader
	var curClass, errStr string

	for {
		rd = bufio.NewReader(conn)
		cmd, args, err := parseCmd(rd)
		if err != nil {
			break
		}
		if cmd == "" {
			continue
		}
		log.Printf("class: [%s] new cmd [%s] - args: %s", curClass, cmd, args)
		actionFunc, ok := actionMap[cmd]
		if !ok {
			conn.Write([]byte(usage))
			continue
		}
		retBytes, err := actionFunc(curClass, args)
		if err != nil {
			errStr = fmt.Sprintf("  class: [%s] [%s] error: %s\n", curClass, cmd, err)
			log.Printf(errStr)
			conn.Write([]byte(errStr))
			continue
		}
		if retBytes != nil {
			conn.Write(retBytes)
		}
		if cmd == "create" || cmd == "select" {
			curClass = args[0]
		}
	}
	conn.Close()
}

func parseCmd(rd *bufio.Reader) (string, []string, error) {
	line, err := rd.ReadString('\n')
	if err != nil {
		return "", nil, err
	}
	argsAll := strings.Fields(strings.TrimSpace(line))
	if len(argsAll) == 0 {
		return "", nil, nil
	} else if len(argsAll) == 1 {
		return argsAll[0], nil, nil
	}
	return argsAll[0], argsAll[1:], nil
}

// func exit(_ string, args []string) ([]byte, error) {
// 	os.Exit(0)
// 	return nil, nil
// }
