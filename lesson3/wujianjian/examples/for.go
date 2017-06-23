package main

import (
	"fmt"
	"os"
)

func main() {
	// for循环的第一种形式
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//for循环的第二种形式,类似于while循环
	i := 5
	for i < 7 {
		fmt.Println(i)
		i = i + 1
	}

	//for循环的第三种形式,等价于while true
	i = 8
	for {
		i = i + 1
		fmt.Println(i)
		if i > 10 {
			break
		}
	}

	//for循环的第四种形式,for ... range
	// i:是下标,arg:是值
	//通过切片来通知下标
	s := "hello中文"

	for i, arg := range s {
		//fmt.Println(i, arg)
		fmt.Printf("%d %c\n", i, arg)
	}

	for _, arg := range os.Args {
		fmt.Println(arg)
	}

}
