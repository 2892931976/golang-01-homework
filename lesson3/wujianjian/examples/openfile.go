package main

import (
	"log"
	"os"
)

func main() {
	//创建文件,不存在创建，存在即清空
	//f, err := os.Create("a.txt")

	// os.O_TRUNC：干净写
	//f, err := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("hello\n")
	//从什么位置开始写,顺序写入
	f.Seek(1, os.SEEK_SET)
	f.WriteString("$$")
	f.Close()
}
