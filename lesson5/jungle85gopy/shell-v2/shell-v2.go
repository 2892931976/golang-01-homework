package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// global var. map makes it easy to delete key
// key : from 1 to CMD_NUM
var cmdMap = make(map[int]*exec.Cmd)
var wrPipe = make(map[int]*io.PipeWriter)

// clear all key of all global Maps
func clearMaps() {
	cmdLen, pipeLen := len(cmdMap), len(wrPipe)
	for i := 1; i <= cmdLen; i++ {
		delete(cmdMap, i)
	}
	for i := 1; i <= pipeLen; i++ {
		delete(wrPipe, i)
	}
	// fmt.Println("cmd:", cmdMap, "\nrd:", rdPipe, "\nwr:", wrPipe)
}

// parse cmd from Scanner, write cmd in global map and return cmd len
func parseCmd(rd *bufio.Scanner) int {
	clearMaps()
	// read from stdin
	if !rd.Scan() {
		os.Exit(0)
	}
	line := rd.Text()
	if len(line) == 0 {
		return 0
	}
	for idx, val := range strings.Split(line, "|") {
		cmd := strings.Fields(val)
		if len(cmd) == 0 {
			return 0
		}
		cmdMap[idx+1] = exec.Command(cmd[0], cmd[1:]...)
	}
	return len(cmdMap)
}

func preparePipe(pipeLen int) {
	for i := 1; i <= pipeLen; i++ {
		rd, wr := io.Pipe()
		wrPipe[i] = wr
		cmdMap[i].Stdout = wr
		cmdMap[i+1].Stdin = rd
	}
	cmdMap[pipeLen+1].Stdout = os.Stdout
	cmdMap[pipeLen+1].Stderr = os.Stderr
}

func runCmd(start int) error {
	// run cmd[start] as single
	if cmdMap[start].Process == nil {
		// Process stores the info about a process
		err := cmdMap[start].Start()
		if err != nil {
			return err
		}
	}
	restLen := len(cmdMap) - start
	if restLen >= 1 {
		//whien restLen>=1, cmd[r+1] and pipe[r] are couples.
		// recursive process couples.
		newStart := start + 1
		err := cmdMap[newStart].Start()
		if err != nil {
			return err
		}
		defer func() {
			// err is the return of cmdMap[start].Wait()
			if err == nil {
				// close pipe while writer exit normally
				wrPipe[start].Close()
				runCmd(newStart)
			}
		}()

	}
	return cmdMap[start].Wait()
}

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[jungle@%s]$ ", host)
	ioScan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		cmdLen := parseCmd(ioScan)
		if cmdLen == 0 {
			continue
		}
		preparePipe(cmdLen - 1)
		if err := runCmd(1); err != nil {
			fmt.Println(err)
		}
	}
}
