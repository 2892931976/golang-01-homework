package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//选择读取方式
	//方法1：裸读取，很少使用,按块读取,速度慢
	buf := make([]byte, 4096)
	n, err := f.Read(buf)
	buf[:n]

	//方法2：加上buffer读取，很高效
	r := bufio.NewReader(f)
	r.Read(buf)

	//方法3：按行读取,按分隔符读取
	r1 := bufio.NewScanner(f)

	//方法4：小文件一次性读取
	f = ioutil.ReadFile("a.txt")
	ioutil.ReadAll(f)

	//方法5：神器，类文件操作:内存、管道
	io.Copy
}
