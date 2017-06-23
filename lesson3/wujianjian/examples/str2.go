package main

import "fmt"

//小写转换为大写函数
func toUpper(s string) string {
	array := []byte(s)
	for i := 0; i < len(array); i++ {
		if 'a' <= array[i] && array[i] <= 'z' {
			//array[i] = array[i] - ('h' - 'H')
			array[i] = array[i] - 32
		}
	}
	return string(array)
}

func main() {
	//相加
	s1 := "hello" + "world"

	s2 := "helloworld"

	//判断字符串相等
	if s1 == s2 {
		fmt.Println("equal")
	}
	//字符串长度
	fmt.Println(0, len(s1)-1)

	//取字符,byte = uint8,是uint的别名
	var c1 byte
	c1 = s1[0]

	fmt.Printf("%d %c\n", c1, c1)

	//切片,默认是左闭右开
	s3 := s1[0:3]
	s4 := s1[:]
	fmt.Println(s1, c1, s3, s4)

	var b byte
	for b = 0; b < 178; b++ {
		fmt.Printf("%d %c\n", b, b)
	}

	//16进制数,a可以大写也可以小写
	fmt.Println(0xa)
	fmt.Println(0xA)

	// ""双引号是字符串,''单引号是字符
	//字符串本身不可修改,可以通过[]byte相互转换进行修改
	array := []byte(s1)
	fmt.Println(array)
	array[0] = 'H' //也可以写成72,结果是一样的
	s1 = string(array)
	fmt.Println(s1)

	fmt.Println(string('a' + ('H' - 'h')))

	fmt.Println(toUpper("hello"))
}
