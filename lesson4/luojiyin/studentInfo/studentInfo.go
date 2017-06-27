package main

import (
	"fmt"
	"io/ioutil"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	//var cmd string
	var name string
	var id int
	//var line string
	//f := bufio.NewReader(os.Stdin)

	for {

		fmt.Print(">")
		fmt.Scanln(&name, &id)
		if name == "stop" {
			break
		}

		if name == "list" {
			dat, err := ioutil.ReadFile("info.txt")
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", dat)
		}
		fmt.Println(name, id)
	}

}
