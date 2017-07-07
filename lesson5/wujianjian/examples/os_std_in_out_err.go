package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Stdin)
	fmt.Println(os.Stdout)
	fmt.Println(os.Stderr)
}
