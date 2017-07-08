package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

// global var. map makes it easy to delete key
// key : from 1 to NUM
var cmdMap = make(map[int]*exec.Cmd)
var rdPipe = make(map[int]*io.PipeReader)
var wrPipe = make(map[int]*io.PipeWriter)

// clear all key of all global Maps
func clearMaps() {
	cmdLen, pipeLen := len(cmdMap), len(rdPipe)
	for i := 1; i <= cmdLen; i++ {
		delete(cmdMap, i)
	}
	for i := 1; i <= pipeLen; i++ {
		delete(rdPipe, i)
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
		cmdMap[idx+1] = exec.Command(cmd[0], cmd[1:]...)
		// fmt.Println("\t", idx+1, "cmd", cmdMap[idx+1])
	}
	return len(cmdMap)
}

func preparePipe(pipeLen int) {
	for i := 1; i <= pipeLen; i++ {
		// fmt.Println("-- create pipe:", i)
		rd, wr := io.Pipe()
		rdPipe[i] = rd
		wrPipe[i] = wr
	}
	// set cmd i(1...) stdout of cmd to pipe wr
	for i := 1; i <= pipeLen; i++ {
		cmdMap[i].Stdout = wrPipe[i]
	}
	// set cmd i(2...) stdin of cmd from pipe rd
	for i := 1; i <= pipeLen; i++ {
		cmdMap[i+1].Stdin = rdPipe[i]
	}
	// very important for show result
	cmdMap[pipeLen+1].Stdout = os.Stdout
}

func runCmd(cmdLen int) {
	for i := 1; i <= cmdLen; i++ {
		cmdMap[i].Start()
	}
	time.Sleep(time.Second * 1)

	for i := 1; i <= cmdLen-1; i++ {
		wrPipe[i].Close()
	}
	for i := 1; i <= cmdLen; i++ {
		cmdMap[i].Wait()
	}
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
		runCmd(cmdLen)
	}
}
