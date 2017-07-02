package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func execFunc(s string) error {
	args := strings.Fields(s)
	if len(args) == 0 {
		return errors.New("command error")
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func execPipeFunc(s string) error {
	args := strings.Split(s, "|")
	var cmdSlice []*exec.Cmd
	for _, v := range args {
		cmd := strings.Fields(v)
		cmdSlice = append(cmdSlice, exec.Command(cmd[0], cmd[1:]...))
	}
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
	prompt := fmt.Sprintf("[test@%s]$ ", host)
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

		n := strings.Index(line, "|")
		if n == -1 {
			err := execFunc(line)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := execPipeFunc(line)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
