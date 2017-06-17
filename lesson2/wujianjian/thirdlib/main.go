package main

import (
	"fmt"
	//"github.com/51reboot/golib"
	golib2 "github.com/DragonWujj/golib"
	"github.com/icexin/golib"
)

func main() {
	var res int
	res = golib.Add(1, 2)
	fmt.Println(res)
	fmt.Println(golib2.Add(2, 3))
}
