package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	//"reflect"
	"strconv"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	//var cmd string
	var action string
	var name string
	var id int
	//var line string
	//f := bufio.NewReader(os.Stdin)

	for {

		fmt.Print(">")
		action = ""
		name = ""
		id = 0
		fmt.Scanln(&action, &name, &id)
		if action == "stop" {
			break
		}

		if action == "list" {
			dat, err := ioutil.ReadFile("info.txt")
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", dat)
			//fmt.Printf("%q\n", strings.Fields(string(dat)))
		}

		if action == "add" && name != "" && id != 0 {
			IsNew := true
			fmt.Println(action, name, id)
			file, err := os.Open("info.txt")
			if err != nil {
				panic(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
				//fmt.Println(reflect.TypeOf(scanner.Text()))
				temp := strings.Fields(scanner.Text())
				if temp[0] == name {
					fmt.Println("No  it has")
					IsNew = false
					break
				}
				if err := scanner.Err(); err != nil {
					panic(err)
				}
			}
			if IsNew {
				f, err := os.OpenFile("info.txt", os.O_APPEND|os.O_WRONLY, 0600)
				if err != nil {
					panic(err)
				}
				defer f.Close()
				fmt.Println("id---------")
				//fmt.Println(id)
				text := name + " " + strconv.Itoa(id) + "\n"
				//fmt.Println(text)
				if _, err = f.WriteString(text); err != nil {
					panic(err)
				}
			}
		}

	}
}
