package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello jcui \n")
	//测试打印当前时间,结果发现格式化系统的layout必须指定如下格式: 2006-01-02 15:04:05.999999999 -0700 MST
	fmt.Fprintf(w, time.Now().Format("2006-01-02 15:04:05 MST"))
	fmt.Println(time.Now().Format("2007-01-02 15:04:05 MST"))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
