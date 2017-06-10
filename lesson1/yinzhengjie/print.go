import "fmt"

func main()  {
var x int  //定义一个数字类型的变量
var y int
x = 100
y = 200
swap(&x,&y)  //将x.y的值取出来当做位置参数传给swap函数，分别传给了p和q
fmt.Print("X==>>>",x,"\n","Y==>>>",y)
}

func swap(p *int,q *int)  {  //该函数定义要求传入2个参数，p和q
var t int  //定义一个数字类型的变量
t = *p  //将p的值传给t，此时p的值为100，故此时t和p的值均为100
*p = *q  //将q的值传给p，此时q的值为200,故此时p的值有100变为了200
*q = t  //将t的值传给q，此刻q的值和t的值均为100，
}
