package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var n int
	var f float32
	n = 10
	//f = float32(n / 3) // 赋值左右的类型应该一致
	f = float32(n) / 3 // 赋值左右的类型应该一致
	fmt.Println(f * 3)
	n = int(f * 10)

	fmt.Println(f, n)

	//大数转小数会发生阶段，小数到大数没有问题
	var n1 int64
	n1 = 1024004 // 1024127|1024129
	var n2 int8
	n2 = int8(n1)
	fmt.Println(n1, n2)

	//整型和字符串转换
	var s string
	// Itoa:整数转换为字符串；Atoi:字符串转整数
	s = strconv.Itoa(n) // string(n),当做字符是有问题的
	fmt.Println(s)

	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(n)

	//FormatInt,FormatUint 2个函数
	//随机数,一定要初始化种子
	var x int64
	rand.Seed(time.Now().Unix())
	x = rand.Int63()
	s = strconv.FormatInt(x, 36) // 36表示进制数,可以为10
	fmt.Println(s)
}
