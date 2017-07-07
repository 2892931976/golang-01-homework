package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

/*

命令1：ls -l | grep go
命令2：cat shell.go | grep main

*/

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[wujianjian@%s]$ ", host)
	r := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		if !r.Scan() {
			break
		}
		line := r.Text()
		if len(line) == 0 {
			continue
		}
		cmds := strings.Split(line, "|")
		s1 := strings.Fields(cmds[0])
		s2 := strings.Fields(cmds[1])
		//fmt.Println(s1)
		//fmt.Println(s2)

		r, w := io.Pipe()
		cmd1 := exec.Command(s1[0], s1[1:]...)
		cmd2 := exec.Command(s2[0], s2[1:]...)
		cmd1.Stdin = os.Stdin
		cmd1.Stdout = w
		cmd2.Stdin = r
		cmd2.Stdout = os.Stdout

		cmd1.Start()
		cmd2.Start()
		cmd1.Wait()
		cmd2.Wait()
	}
}
