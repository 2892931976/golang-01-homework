package main

import (
	"os"
	"net/url"
	"log"
	"fmt"
)

func main() {

	 s := os.Args[1]
	 u, err := url.Parse(s)
	if err != nil {
		  log.Fatal(err)

	}

	fmt.Println("scheme", u.Scheme) //请求协议
	fmt.Println("host", u.Host) //请求主机
	fmt.Println("path", u.Path) //路径
	fmt.Println("queryString", u.RawQuery) //请求信息
	fmt.Println("user", u.User)   //用户验证
	fmt.Println("xx", u.Fragment) //锚点

}
