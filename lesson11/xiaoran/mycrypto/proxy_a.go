package main

import (
	"crypto/md5"
	"crypto/rc4"
	"flag"
	"io"
	"log"
	"net"
	"sync"
)

var (
	target = flag.String("target", "127.0.0.1:8080", "target host")
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
	i, err := w.w.Write(buf)
	return i, err
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
	n, err := r.r.Read(b)
	buf := b[:n]
	r.cipher.XORKeyStream(buf, buf)
	return n, err
}

func handleConn(conn net.Conn) {
	//建立到目标服务器的连接
	var remote net.Conn
	var err error
	remote, err = net.Dial("tcp", *target)
	if err != nil {
		log.Print(err)
		conn.Close()
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	//go 读取（conn)的数据，发送到remote，直到conn的EOF，关闭remote
	go func() {
		defer wg.Done()
		w := NewCryptoWriter(remote, "123456")
		io.Copy(w, conn)
		remote.Close()
	}()

	//go 读取remote的数据，发送到客户端（conn）,直到remote的EOF，关闭conn
	go func() {
		defer wg.Done()
		r := NewCryptoReader(remote, "123456")
		io.Copy(conn, r)
		conn.Close()
	}()

	wg.Wait()
	//等待连接关闭
}

func main() {
	flag.Parse()
	addr := ":7070"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
