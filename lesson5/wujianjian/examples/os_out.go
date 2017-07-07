package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	f, err := os.Create("ls.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = f
	//等价于cmd.Run()
	cmd.Start()
	cmd.Wait()
}
