/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

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
	target = flag.String("target", "www.baidu.com:80", "target host")
)

type CryptoWriter struct {
	w      io.Writer
	cipher *rc4.Cipher
}

func NewCryptoWriter(w io.Writer, key string) io.Writer { /*实现了io.writer的接口，当你实现了某个write的方法，那你就实现了write接口。
	NewCryptoWriter是构造函数，相当于python中的__init__*/
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher([]byte(md5sum[:]))
	if err != nil {
		log.Fatal(err)
	}
	return &CryptoWriter{
		w:      w,
		cipher: cipher,
	}
}

func (w *CryptoWriter) Write(b []byte) (int, error) {
	buf := make([]byte, len(b))
	w.cipher.XORKeyStream(buf, b) //进行加密
	num, err := w.w.Write(buf)    //将加密的后的数据写入到w中去。
	return num, err
}

type CryptoReader struct {
	r      io.Reader
	cipher *rc4.Cipher
}

func NewCryptoReader(r io.Reader, key string) io.Reader {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher([]byte(md5sum[:]))
	if err != nil {
		log.Fatal(err)
	}
	return &CryptoReader{
		r:      r,
		cipher: cipher,
	}
}

func (r *CryptoReader) Read(b []byte) (int, error) {
	num, err := r.r.Read(b) //n和b的长度不一定相等。
	buf := b[:num]
	r.cipher.XORKeyStream(buf, buf)
	return num, err
}

func handle_Conn(conn net.Conn) {
	var (
		remote net.Conn //定义远端的服务器连接。
		err    error
	)
	remote, err = net.Dial("tcp", *target) //建立到目标服务器的连接。
	if err != nil {
		log.Print(err)
		conn.Close()
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		w := NewCryptoWriter(conn, "123456") //在将数据发送给加密隧道的另一端的时候，需要把数据进行加密。
		io.Copy(w, remote)                   //读取原地址请求（conn），然后将读取到的数据发送给w。
		remote.Close()
	}()
	go func() {
		defer wg.Done()
		r := NewCryptoReader(conn, "123456")
		io.Copy(remote, r) //与上面相反，就是讲目标主机的数据返回给客户端。
		conn.Close()
	}()
	wg.Wait()
}

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, _ := l.Accept()
		go handle_Conn(conn)
	}

}

/*
	1>.启动proxy-A,proxy-B,socks5服务器并监听端口；
	2>.浏览器需要指定proxy-A的地址和端口；
	3>.proxy-A进行拨号至proxy-A监听的端口，和proxy-B建立链接;
	4>.proxy-B进行拨号，至目标服务器，然后服务器又会将数据返回回来，然后就是一个想反的方向。
*/
