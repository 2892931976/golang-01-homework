/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import "fmt"

var x int = 200

func localFunc() {
	fmt.Println(x)
}

func main() {
	x := 1
	localFunc()    //打印数字200，因为该函数中和他同级的是全局变量x = 200
	fmt.Println(x) //打印数字1,他在寻找变量的时候回从同级找，从缩进来看，都是在main函数内部的作用于，有x = 1的变量，故打印出来了
	if true {
		x := 100       //给x变量赋初值，也就是我们说的短变量，要注意的是它仅在局部变量使用,在全局变量使用会报错。
		fmt.Println(x) //打印数字100，因为它还是会从它同级开始找，从if语句开始，然后再到main函数，再到全局变量一次向外查询。
	}

	localFunc()    //打印数字200，因为该函数中和他同级的是全局变量x = 200，道理一样，这个localFunc函数回去它同级作用域（对于该函数而已就是全局变量中）找，找到就打印，找不到就报错。
	fmt.Println(x) //打印数字1，还是在它的同级作用域查找，其同级作用域就是缩进相同的变量或函数或者条件表达式等等。
}
