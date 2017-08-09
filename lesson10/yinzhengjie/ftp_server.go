/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"fmt"
	"path/filepath"
)

var (
	cmd       string
	file_name string
)

func handel_conn(conn net.Conn) {
	r := bufio.NewReader(conn)      //“r”会进行顺序读取，即流式处理，当它读过的内容就无法继续往回读取，知道遇到EOF结束标识符才会终止。
	line, err := r.ReadString('\n') //循环读取“r”的内容，遇到“\n”换行符就中止。
	if err != nil {
		log.Print(err)
	}
	fmt.Println(line)
	line = strings.TrimSpace(line) //去掉换行符
	fields := strings.Fields(line) //按照空格分隔字符串。
	cmd := fields[0] //定义用户输入的
	file_name := fields[1]
	switch cmd {
	case "GET","get":
		f, err := os.Open(file_name)
		if err != nil {
			log.Print(err)
			return
		}
		/*
		   处理文件的方式一：
				buf := make([]byte,4096) //定义一个切片大小，用于按块读取数据。
				for{
					n,err := f.Read(buf)
					if err == io.EOF {  //指定循环结束条件，当读取到文件的结束标识符是中断循环。
						break
					}
					conn.Write(buf[:n]) //讲读取到的内容发送给客户端。
				}

		   处理文件的方式二：
				buf,err := ioutil.ReadAll(f) //将文件全部读取出来，但仅仅适合读取小文件。不推荐使用、
				if err != nil {
					log.Print(err)
					return
				}
				conn.Write(buf)
		*/
		io.Copy(conn, f) //其运行机制就是循环读取按块读取“f”中的内容然后讲读取的数据传递给“conn”,此种方式最为高效率。
		f.Close()
		fmt.Println("读取完毕！")
	case  "STORE","store" :
		os.MkdirAll(filepath.Dir(file_name), 0755) //获取name的目录，并在服务器中创建出来。
		f1, err := os.Create(file_name)
		if err != nil {
			log.Print(err)
			return
		}
		io.Copy(f1, r) //讲链接读取到的内容写入到刚刚创建的文件中。
		f1.Close()
		//conn.Close()
		fmt.Println("STORE命令执行完毕!")
	}
}

func main() {
	addr := "0.0.0.0:8080" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go handel_conn(conn)
	}
}
