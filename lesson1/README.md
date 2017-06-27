# 一. 完善http.go和c10k.go

要求：

1. 编译通过

2. 编译成windows, linux和mac三个平台的文件，有条件的可以运行观察结果。二进制不用提交

# 二. 分析课堂上的代码，说明原因
```go
func main() {
    var x int
    var y int
    x = 1
    y = 2
    swap(&x, &y)
    fmt.Println("x=", x, "y=", y)
}

func swap(p *int, q *int) {
    var t int
    t = *p
    *p = *q
    *q = t
}
```
