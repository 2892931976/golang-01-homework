package main



import (
	"bufio"
	//"errors"
	"fmt"
	"os"
	"strings"
	"strconv"
	"encoding/json"
	"io/ioutil"
)
type Student struct{
	Id int
	Name string
}


var infos = make(map[string]Student)

func add(args []string)error{
	//fmt.Println("call add")
	//fmt.Println("args", args)
	name := args[0]
	id,err := strconv.Atoi(args[1])
	if err != nil {
		err := fmt.Errorf("%s ,%s", "id type is not int", err)
		return err
	}

	infos[name] = Student{Id:id,Name:name}
	fmt.Println(infos)
	return nil
}

func list(args []string)error{
	for _,v := range infos{
		fmt.Println(v.Name,v.Id)
	}
	return nil
	//return errors.New("unimplemention")
}


func save(args []string) error{
		file := args[0]
		f,err := json.Marshal(infos)
		if err != nil{
			err := fmt.Errorf("%s,%s","load file",err)
			return err
		}
		ioutil.WriteFile(file, f, 0400)
		fmt.Println("save ok")
		return nil
}





func load(args []string) error{
	file := args[0]
	f, err := ioutil.ReadFile(file)
	if err != nil {
		err := fmt.Errorf("%s ,%s", "load faile", err)
		return err
	}
	json.Unmarshal(f, &infos)
	fmt.Println("load ok")
	return nil


}


func main(){
	actionmap := map[string]func([]string)error{
		"add":add,
		"list":list,
		"save":save,
		"load":load,
	}

	f := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("> ")
		line,_ := f.ReadString('\n')

		args := strings.Fields(line)
		if len(args) == 0{
			continue
		}
		cmd := args[0]
		args = args[1:]

		actionfunc := actionmap[cmd]
		if actionfunc == nil {
			fmt.Println("bad cmd ", cmd)
			continue
		}
		err := actionfunc(args)
		if err != nil {
			fmt.Println("execute action %s error:%s\n", cmd, err)
			continue
		}
	}

}