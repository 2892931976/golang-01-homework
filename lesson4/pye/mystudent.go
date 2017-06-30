package main

import (
	"bufio"
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

func main() {
	var cmd string
	var name string
	var filename string
	var id int
	var line string
	f := bufio.NewReader(os.Stdin)

	var students map[string]Student
	students = make(map[string]Student)

	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			for _, info := range students {
				fmt.Println(info.Name, info.Id)
			}

		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			_, ok := students[name]

			if ok {
				fmt.Printf("%v has benn added \n", name)
				break
			} else {
				students[name] = Student{Id: id, Name: name}
			}
			fmt.Printf("%v added.\n", name)

		case "save":
			fmt.Sscan(line, &cmd, &filename)
			_m2json, _ := json.Marshal(students)
			outputFile, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			defer outputFile.Close()
			outputWriter := bufio.NewWriter(outputFile)
			outputWriter.WriteString(string(_m2json))
			outputWriter.Flush()

			fmt.Printf("saved \n")

		case "load":
			fmt.Sscan(line, &cmd, &filename)
			f, err := ioutil.ReadFile(filename)

			if err == nil {
				_json := json.Unmarshal(f, &students)
				if _json != nil {
					fmt.Println(_json)
				}
			} else {
				log.Fatal(err)
			}

			fmt.Printf("loaded %v \n", filename)

		case "quit":
			os.Exit(0)
		}
	}
}
