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
	//"io/ioutil"
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
	fields := strings.Fields(line) //按照空格分隔字符串
	if len(fields) != 2 {
		conn.Write([]byte("bad input!"))
		return
	}
	cmd := fields[0] //定义用户输入的
	name := fields[1]
	if cmd == "GET" {
		f, err := os.Open(name)
		if err != nil {
			log.Print(err)
			return
		}
		defer f.Close()
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

	} else if cmd == "STORE" {
		os.MkdirAll(filepath.Dir(name), 0755) //获取name的目录，并在服务器中创建出来。
		f, err := os.Create(name)
		if err != nil {
			log.Print(err)
			return
		}
		io.Copy(f, r) //讲链接读取到的内容写入到刚刚创建的文件中。
		defer f.Close()
		/*
			1.从"r"读取文件内容直到err为io.EOF
			2.创建name文件
			3.向文件写入数据
			4.往conn写入OK
			5.关闭连接和文件。
		*/
		conn.Write([]byte("这是上传文件的方法"))
	}

	var content []byte
	//读取文件内容到conntent
	conn.Write(content)
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
