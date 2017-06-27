package main

import "fmt"

func main() {
	/*
		          \n： 换行
			  \t： 制表符
			  \\： 转义\
			  \b：
			  \r：
			  \"：打印双引号
			  \a：响铃
	*/
	str1 := "hello\\\"world\a"
	doc := `
即使换行也不影响
可以验证一下
类似python的'''
`
	fmt.Println(str1, doc)
}
