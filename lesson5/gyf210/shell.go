package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func argsFunc(s string) []*exec.Cmd {
	var cmdSlice []*exec.Cmd
	n := strings.Index(s, "|")
	if n == -1 {
		args := strings.Fields(s)
		cmdSlice = append(cmdSlice, exec.Command(args[0], args[1:]...))
		return cmdSlice
	}
	args := strings.Split(s, "|")
	for _, v := range args {
		cmd := strings.Fields(v)
		cmdSlice = append(cmdSlice, exec.Command(cmd[0], cmd[1:]...))
	}
	return cmdSlice
}

func execFunc(cmdSlice []*exec.Cmd) error {
	pipeSlice := make([]*io.PipeWriter, len(cmdSlice)-1)
	i := 0
	for ; i < len(cmdSlice)-1; i++ {
		r, w := io.Pipe()
		cmdSlice[i].Stdout = w
		cmdSlice[i+1].Stdin = r
		pipeSlice[i] = w
	}
	cmdSlice[i].Stdout = os.Stdout
	cmdSlice[i].Stderr = os.Stderr

	err := callFunc(cmdSlice, pipeSlice)
	if err != nil {
		return err
	}
	return nil
}

func callFunc(cmdSlice []*exec.Cmd, pipeSlice []*io.PipeWriter) error {
	if cmdSlice[0].Process == nil {
		err := cmdSlice[0].Start()
		if err != nil {
			return err
		}
	}

	if len(cmdSlice) > 1 {
		err := cmdSlice[1].Start()
		if err != nil {
			return err
		}
		defer func() {
			if err == nil {
				pipeSlice[0].Close()
				err = callFunc(cmdSlice[1:], pipeSlice[1:])
			}
		}()
	}
	return cmdSlice[0].Wait()
}

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[TEST@%s]$ ", host)
	r := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)

		if !r.Scan() {
			break
		}

		line := r.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		cmdSlice := argsFunc(line)
		err := execFunc(cmdSlice)
		if err != nil {
			fmt.Println(err)
		}
	}
}
