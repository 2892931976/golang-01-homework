package main

import (
	"bufio"
	"crypto/md5"
	"crypto/rc4"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

func Crypto(w io.Writer, r io.Reader, key string) {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil {
		log.Print("Crypto error:", err)
		return
	}
	buf := make([]byte, 4096)

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		cipher.XORKeyStream(buf[:n], buf[:n])
		w.Write(buf)
	}

}

func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	//加密的过程
	key := GetRandomString(10)
	r_buf := bufio.NewReader(conn)
	fmt.Println(key, r_buf)

	//向服务器端发送数据
	server, err := net.Dial("tcp", "fps6.uc.ppweb.com.cn:7777")
	if err != nil {
		log.Print("Error:", err)
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	//go 接收客户端的数据,发送到remote,直到conn的EOF
	go func() {
		defer wg.Done()
		io.Copy(server, conn)
		server.Close()
	}()
	//解密的过程

	//go 接收remote的数据,发送到客户端,直到remote的EOF
	go func() {
		defer wg.Done()
		io.Copy(conn, server)
		conn.Close()
	}()
	//等待两个协程结束
	wg.Wait()

}

func main() {
	//建立监听
	addr := ":7777"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	//接收新的连接
	for {
		//accept new conection
		conn, err := listener.Accept()
		log.Print(conn.RemoteAddr())
		if err != nil {
			log.Print(err)
		}
		// 参考页面 http://www.jianshu.com/p/172810a70fad
		go handleConn(conn)
	}
}
