package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		args := strings.Split(line, "|")
		for _, arg := range args {
			fmt.Println(arg)
			word := strings.Fields(arg)
			cmd := exec.Command(word[0], word[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
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
