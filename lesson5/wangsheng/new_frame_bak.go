package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Student struct {
	Id int
	Name string

}

var stumap = make(map[string]Student)



func add(args []string) error{

	fmt.Println("call add")
	//取得
	name := args[0]
	id,err := strconv.Atoi(args[1])
	if err != nil {
		err := fmt.Errorf("%s ,%s", "id type is not int", err)
		return err
	}
	fmt.Print(name)
	fmt.Print(id)

	return nil
}


func main(){

	//funcmap := map[string]func([]string) error{
	//	"list":list,
	//	"add":add,
	//}
	funcmap := map[string]func([]string) error{
		"add":    add,
	}

	f := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, _ :=  f.ReadString('\n')
		//fmt.Print(line)
		args := strings.Fields(line)
		//变成了slice
		//fmt.Print(args)

		cmd := args[0]
		//fmt.Print(cmd)
		args = args[1:]
		//fmt.Print(args)
		
		//action := funcmap(cmd)
		//错误的写法

		action := funcmap[cmd]
		fmt.Println(action)
		if action == nil {
			fmt.Println("bad cmd:", cmd)
			continue
		}

	}


}