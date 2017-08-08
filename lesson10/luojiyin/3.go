package main

import  {
	"flag"
	"net"
}

var (
	target = flag.String("target", "www.baidu.com:80", "target host")
)
func handleConn(conn net.Conn) {
	var remote net.Conn
	var err  error
	// go 接收 客户端的数据 发送到remote
	//go 接收 remote数据    发送到客户端端

	remote , err := net.Dial("tcp", *target)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		io.Copy(remote, conn)
		conn.Close()
		return
 	}()

	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
		remote.Close()
	}()
	wg.Wait()
}

func main() {
	l ,err := net.Listen("tcp", 8080)
	for {
		handleConn(l)
	}
}
