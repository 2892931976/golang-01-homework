package main

import  "fmt"

func main() {
    var x int 
    var y int
    x = 1
    y = 2
    swap = (&x, &y)// 获取 变量x，y的地址
    fmt.Println("x=", x, "y=", y)
}

func swap( p *int, q *int){
// p, q 分别为整型指针
    var t int
    //初始化一个整数， 作为交换整数，保存临时值
    t = *p // 指针p指向的值，赋值给 t
    *p = *q// 指针q指向的值，赋值给 指针p指向的变量
    *q = t // 把变量t 的值，赋值给 指针q指向的变量 
}    
