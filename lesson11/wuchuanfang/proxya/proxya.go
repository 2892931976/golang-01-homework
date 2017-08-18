package main

import (
	"crypto/md5"
	"crypto/rc4"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type CryptoWriter struct {
	w      io.Writer
	cipher *rc4.Cipher
}

func NewCryptoWriter(w io.Writer, key string) io.Writer {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil {
		panic(err)
	}
	return &CryptoWriter{
		w:      w,
		cipher: cipher,
	}
}

//把b里面的数据进行加密，之后写入到w.w里面
//调用w.w.Write进行写入
func (w *CryptoWriter) Write(b []byte) (int, error) {
	buf := make([]byte, len(b))

	w.cipher.XORKeyStream(buf, b)
	return w.w.Write(buf)

}

type CryptoReader struct {
	r      io.Reader
	cipher *rc4.Cipher
}

func NewCryptoReader(r io.Reader, key string) io.Reader {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil {
		panic(err)
	}
	return &CryptoReader{
		r:      r,
		cipher: cipher,
	}

}

func (r *CryptoReader) Read(b []byte) (int, error) {
	//buf := make([]byte, 1024)
	n, err := r.r.Read(b)

	if err != nil {
		log.Fatal(err)

	}
	buf := b[:n]
	r.cipher.XORKeyStream(buf, buf)
	return n, err

}

var (
	target = flag.String("target", "127.0.0.1:8080", "target host")
)

func handleConn(conn net.Conn) {
	//建立到目标服务器的连接
	//defer conn.Close()
	//var remote net.Conn
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println(*target)
	remote, err := net.Dial("tcp", *target)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer wg.Done()
		w := NewCryptoWriter(remote, "123456")
		io.Copy(w, conn)
		remote.Close()
	}()
	go func() {
		defer wg.Done()
		r := NewCryptoReader(remote, "123456")
		io.Copy(conn, r)
		conn.Close()
	}()
	wg.Wait()
	//go 接收客户端的数据发送到remote
	//go 接收remote的数据，发送到客户端
	//等待连接关闭
}

func main() {
	//建立监听
	addr := "127.0.0.1:7070"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	//defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(conn)
	}

}
