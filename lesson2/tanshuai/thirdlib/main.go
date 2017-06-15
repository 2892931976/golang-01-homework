package main

import (
	"fmt"
	//go调用别人的类库方法
	"github.com/icexin/golib"
)

func main() {
	fmt.Println(golib.Add(7, 2))
}
