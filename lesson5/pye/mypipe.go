package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("%s >>> ", host)
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

		if !strings.ContainsAny(line, "|") {
			args := strings.Fields(line)
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}

		} else {
			args := strings.Split(line, "|")

			c1 := strings.Fields(args[0])
			c2 := strings.Fields(args[1])

			pCmd := exec.Command(c1[0], c1[1:]...)
			nCmd := exec.Command(c2[0], c2[1:]...)

			reader, writer := io.Pipe()
			var buf bytes.Buffer
			pCmd.Stdout = writer
			nCmd.Stdin = reader
			nCmd.Stdout = &buf
			pCmd.Start()
			nCmd.Start()
			pCmd.Wait()
			writer.Close()
			nCmd.Wait()
			reader.Close()
			io.Copy(os.Stdout, &buf)
		}
	}
}
