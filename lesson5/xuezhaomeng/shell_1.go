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

	prompt := fmt.Sprintf("[xuezhaomeng@%v]$ ", host)
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
		if strings.Contains(line, "|") {
			cmds := strings.Split(line, "|")
			s1 := strings.Fields(cmds[0])
			s2 := strings.Fields(cmds[1])

			//读取 ,写入
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

			//不匹配的结果的错误提醒也输出,暂时不会解决
			//bug1: 命令1 的结果也进行了输出
		}
		args := strings.Fields(line)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}

	}
}
