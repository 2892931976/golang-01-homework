package main

import (
	"fmt"
	"os"
)

func myreserverv2(s string) string {
	reserv2 := []byte(s)
	for i, j := 0, len(reserv2)-1; i < j; i, j = i+1, j-1 {

		reserv2[i], reserv2[j] = reserv2[j], reserv2[i]

	}

	return string(reserv2)

}
func main() {
	if len(os.Args) != 2 {
		fmt.Println("you should input args 1")

	} else {
		zifu := os.Args[1]
		fmt.Println(myreserverv2(zifu))
	}

}
