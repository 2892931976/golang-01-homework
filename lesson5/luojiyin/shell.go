package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[icexin@%s]$ ", host)
	r := bufio.NewScanner(os.Stdin)
	//r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		if !r.Scan() {
			break
		}
		line := r.Text()
		// line, _ := r.ReadString('\n')
		// line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		//args := strings.Fields(line)
		//tempStdin := os.Stdin
		var tempStdin *os.File
		var temp *os.File
		args := strings.Split(line, "|")
		for i, arg := range args {
			fmt.Println(arg)
			word := strings.Fields(arg)
			cmd := exec.Command(word[0], word[1:]...)
			//tempStdin, _ := cmd.StdoutPipe()
			//fmt.Println(tempStdin.(type))
			temp, _ = cmd.StdoutPipe()
			fmt.Println("type", reflect.TypeOf(tempStdin))
			if i < 0 {
				cmd.Stdin = os.Stdin
			}
			if i > 0 {
				cmd.Stdin = temp
			}

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
			//temp, err := cmd.StdoutPipe()
			fmt.Println("type 23 ", reflect.TypeOf(temp))
		}
		/*cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}*/
	}
}
