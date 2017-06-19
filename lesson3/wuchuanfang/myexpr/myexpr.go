package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//fmt.Println(os.Args)
	if len(os.Args) != 4 {
		fmt.Println("Input Error! Please enter int like '11 +|-|*|/ 12'")
	} else {

		s1 := os.Args[1]
		s2 := os.Args[2]
		s3 := os.Args[3]
		//fmt.Println(s1, s2, s3)

		n1, err1 := strconv.Atoi(s1)
		n3, err2 := strconv.Atoi(s3)
		if (err1 != nil) || (err2 != nil) {
			//fmt.Println(err1, err2)
			fmt.Println("Input Error! Please enter int like '11 +|-|*|/ 12'")
			os.Exit(0)
		}

		//fmt.Println(n1, n3)

		switch s2 {
		case "+":
			fmt.Println(n1 + n3)
		case "-":
			fmt.Println(n1 - n3)
		case "*":
			fmt.Println(n1 * n3)
		case "/":
			fmt.Println(n1 / n3)
		default:
			fmt.Println("Input Error! Please enter int like '11 +|-|*|/ 12'")
			os.Exit(0)
		}

	}
}
