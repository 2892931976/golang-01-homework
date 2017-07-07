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

func main() {
	rd := bufio.NewScanner(os.Stdin)
	if !rd.Scan() {
		os.Exit(0)
	}
	line := rd.Text()
	if len(line) == 0 {
		os.Exit(0)
	}

	tmpArr := strings.Split(line, "|")
	fmt.Println("\n--- start read ---")
	s1 := strings.Fields(tmpArr[0])
	s2 := strings.Fields(tmpArr[1])

	// w | r;  writer | reader
	// each Write to the PipeWriter blocks until it has satisfied
	// Reads from the PipeReader that fully consume the written data.
	rp, wp := io.Pipe()
	cmd1 := exec.Command(s1[0], s1[1:]...)
	cmd2 := exec.Command(s2[0], s2[1:]...)

	cmd1.Stdout = wp
	// stdout to writer
	cmd2.Stdin = rp
	// stdin from reader
	cmd2.Stdout = os.Stdout

	cmd1.Start()
	cmd2.Start()
	fmt.Println("-- after start ...")

	cmd1.Wait()
	// cmd2.Wait()
	// Wait() 不能放在这个位置。严格的说，cmd2.Wait(）不在放在这里。因为：
	// 当cmd1在这里wait是会正常结束。然后停止向wp写数据，但w并没有关闭；
	// 而cmd2在这里wait时，其输入来自rp，而rp在读完数据后并没有关闭，
	// 导致cmd2的Stdin也一直等着，从而cmd2并不会退出，从而Wait()也不结束。
	time.Sleep(time.Second * 1)
	wp.Close()

	cmd2.Wait()
	fmt.Println("-- after wait. exit ...")
}
