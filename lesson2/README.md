# 一、 完成mycat

```go
package main

import (
        "fmt"
        "io/ioutil"
)

func printFile(name string) {
        buf, err := ioutil.ReadFile(name)
        if err != nil {
                fmt.Println(err)
                return
        }
        fmt.Println(string(buf))
}

func main() {
        // 补全缺失的代码完成cat命令
}
```

### 要求:
- 程序放到个人的mycat目录下

### 提示:
- 使用os.Args获取命令行参数，具体代码参照myecho


# 二、判断如下程序的输出，并说明原因

```go
package main

import "fmt"

var x = 200

func localFunc() {
	fmt.Println(x)
}

func main() {
	x := 1

	localFunc()
	fmt.Println(x)
	if true {
		x := 100
		fmt.Println(x)
	}

	localFunc()
	fmt.Println(x)
}

```
