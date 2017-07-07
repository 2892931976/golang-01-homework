package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[wuchf@%s]$ ", host)
	r := bufio.NewScanner(os.Stdin)
	//r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		if !r.Scan() {
			break
		}
		line := r.Text()
		//line, _ := r.ReadString('\n')
		//line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		//fmt.Println(line)
		//args := strings.Fields(line)
		cmds := strings.Split(line, "|")

		//var args []string
		var cmdmap = make(map[int][]string)
		//var cmdss map[int][]string
		for i, s := range cmds {
			cmdmap[i] = strings.Fields(s)
		}
		//fmt.Println(cmdmap[0])
		switch len(cmdmap) {
		case 1:
			cmd1 := exec.Command(cmdmap[0][0], cmdmap[0][1:]...)
			//cmd2 := exec.Command(cmdmap[1][0], cmdmap[1][1:]...)
			//fmt.Println(cmd1)
			cmd1.Stdin = os.Stdin
			cmd1.Stdout = os.Stdout
			cmd1.Stderr = os.Stderr
			err := cmd1.Run()
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			cmd1 := exec.Command(cmdmap[0][0], cmdmap[0][1:]...)
			cmd2 := exec.Command(cmdmap[1][0], cmdmap[1][1:]...)
			r, w := io.Pipe()
			cmd1.Stdin = os.Stdin
			cmd1.Stdout = w
			cmd2.Stdin = r
			cmd2.Stdout = os.Stdout
			cmd2.Stderr = os.Stderr
			err1 := cmd1.Start()
			err2 := cmd2.Start()
			cmd1.Wait()

			if err1 != nil || err2 != nil {
				fmt.Println(err1, err2)
			}

		default:
			fmt.Println(prompt)
		}

	}
}
