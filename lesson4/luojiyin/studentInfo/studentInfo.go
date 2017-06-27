package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
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
			fmt.Printf("%q\n", strings.Fields(string(dat)))
		}

		if action == "add" && name != "" && id != 0 {
			file, err := os.Open("info.txt")
			if err != nil {
				panic(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
				fmt.Println(reflect.TypeOf(scanner.Text()))
				temp := strings.Fields(scanner.Text())
				if temp[0] == name {
					fmt.Println("No  it has")
					break
				}
			}
			if err := scanner.Err(); err != nil {
				panic(err)
			}
		}
	}

}
