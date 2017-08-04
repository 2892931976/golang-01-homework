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
	prompt := fmt.Sprintf("[xiao@%s]$ ", host)
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

		//对输入的命令切割
		cmds := strings.Split(line, "|")
		if len(cmds) == 1 {
			args := strings.Fields(line)
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Start()
			cmd.Wait()
			continue
		}

		s1 := strings.Fields(cmds[0])
		s2 := strings.Fields(cmds[1])

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
	}
}
