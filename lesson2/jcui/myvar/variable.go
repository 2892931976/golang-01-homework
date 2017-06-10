package myvar

import "fmt"
import "net"

func main() {
	// 定义全局变量
	// 每个类型都有默认的初始值-零值(安全)
	var (
		x int
		y float32
		z string
		p *int
		a bool
	)
	// 如下是定义局部变量
	i := 0       //int
	s := "hello" //string
	m, n := 0, 1 //批量初始化

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", s)
	fmt.Printf("%v %v\n", m, n)
}
