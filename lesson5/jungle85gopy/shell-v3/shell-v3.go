// shell-v2/shell-v2.go use global map variable to store cmd and pipe.
// it's not a good idea. local slice is well.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// parse cmd from Scanner
func parseCmd(rd *bufio.Scanner) (cmds []*exec.Cmd) {
	// read from stdin
	if !rd.Scan() {
		os.Exit(0)
	}
	line := rd.Text()
	if len(line) == 0 {
		return nil
	}
	for _, val := range strings.Split(line, "|") {
		cmdStr := strings.Fields(val)
		if len(cmdStr) == 0 {
			return nil
		}
		cmds = append(cmds, exec.Command(cmdStr[0], cmdStr[1:]...))
	}
	return
}

func execByPipe(cmds []*exec.Cmd) error {
	var wrPipe []*io.PipeWriter
	pipeLen := len(cmds) - 1

	for i := 0; i < pipeLen; i++ {
		rd, wr := io.Pipe()
		wrPipe = append(wrPipe, wr)
		cmds[i].Stdout = wr
		cmds[i+1].Stdin = rd
	}
	cmds[pipeLen].Stdout = os.Stdout
	cmds[pipeLen].Stderr = os.Stderr
	return runCmd(cmds, wrPipe)
}

func runCmd(cmds []*exec.Cmd, wrPipe []*io.PipeWriter) error {
	// fmt.Println("-- runCmd() rest len= ", len(cmds))
	// run cmd[start] as single
	if cmds[0].Process == nil {
		// Process stores the info about a process
		if err := cmds[0].Start(); err != nil {
			return err
		}
	}
	if len(cmds) > 1 {
		//while len>1, cmds[r+1] and pipe[r] are couples.
		// recursive process couples.
		err := cmds[1].Start()
		if err != nil {
			return err
		}
		defer func() {
			// err is the return of cmdMap[start].Wait()
			if err == nil {
				// close pipe while writer exit normally
				wrPipe[0].Close()
				runCmd(cmds[1:], wrPipe[1:])
			}
		}()

	}
	return cmds[0].Wait()
}

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[jungle@%s]$ ", host)
	ioScan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		cmds := parseCmd(ioScan)
		if len(cmds) == 0 {
			continue
		}
		err := execByPipe(cmds)
		if err != nil {
			fmt.Println(err)
		}
	}
}
