package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		// int后面的数字是位数，除以8就是字节数
		x  byte   // 1个字节
		x1 int    // 8个字节
		x2 int32  // 4个字节
		x3 int64  // 8个字节
		x4 uint   // 8个字节
		x5 uint32 // 4个字节
		x6 uint64 // 8个字节

		x7 int8  = -128 // 1个字节，范围：-128 ~ 127
		x8 uint8 = 255  // 1个字节，范围：0 ~ 255
	)
	//数据溢出
	x7 = 127
	x7 = x7 + 1 //结果变为: -128，但是还在自己的取值范围之内
	x8 = 255
	x8 = x8 + 1 //结果变为：0，但是还在自己的取值范围之内
	fmt.Println(x, x1, x2, x3, x4, x5, x6, x7, x8)
	//unsafe.Sizeof()计算字节大小
	fmt.Println(unsafe.Sizeof(x), unsafe.Sizeof(x1), unsafe.Sizeof(x2), unsafe.Sizeof(x3), unsafe.Sizeof(x4), unsafe.Sizeof(x5), unsafe.Sizeof(x6), unsafe.Sizeof(x7), unsafe.Sizeof(x8))
}
