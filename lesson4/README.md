一、完成切片旋转

指定切片的一个位置，旋转位置两侧的数据

如切片为`[2, 3, 5, 7, 11]`，位置为2，则翻转后的切片为`[5, 7, 11, 2, 3]`

完成切片旋转的同学可以尝试旋转单次，如`hello world`旋转为`world hello`


二、完成学生信息管理系统

实现如下4个指令

1. `add name id`，添加一个学生的信息，如果name有重复，报错
2. `list`, 列出所有的学生信息
3. `save filename`，保存所有的学生信息到filename指定的文件中
4. `load filename`, 从filename指定的文件中加载学生信息


框架代码如下:

``` go
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	var cmd string
	var name string
	var id int
	var line string
	f := bufio.NewReader(os.Stdin)

	//var students map[string]Student
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Printf("binggan 01\njack 02\n")
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			fmt.Printf("add done.\n")
		}
	}
}
```
