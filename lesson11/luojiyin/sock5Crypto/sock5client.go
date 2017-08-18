package main

import (
	"crypto/md5"
	"crypto/rc4"
	"io"
	"log"
	"net"
	"sync"
)

const key = "123456"

func main() {
	listener, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		lisConn, _ := listener.Accept()
		go handleConn(lisConn)
		log.Println("start connect to proxy server")
	}
}

func handleConn(listenConn net.Conn) error {
	log.Println("star ")
	defer listenConn.Close()

	remoteConn, err := net.Dial("tcp", "139.162.109.162:8022")
	if err != nil {
		log.Fatal(err)
	}
	defer remoteConn.Close()

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		remoteCryWriter := NewCrytoWrite(remoteConn, key)
		io.Copy(remoteCryWriter, listenConn)
	}()
	go func() {
		defer wg.Done()
		remoteCryReader := NewCrytoReader(remoteConn, key)
		io.Copy(listenConn, remoteCryReader)
	}()
	wg.Wait()
	return nil
}

type CryptoWriter struct {
	w      io.Writer
	cipher *rc4.Cipher
}

func NewCrytoWrite(w io.Writer, key string) io.Writer {
	md5Sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5Sum[:])
	if err != nil {
		log.Print(err)
	}
	return &CryptoWriter{
		w:      w,
		cipher: cipher,
	}
}

func (w *CryptoWriter) Write(b []byte) (int, error) {
	buf := make([]byte, len(b))

	w.cipher.XORKeyStream(buf, b)
	w.w.Write(buf)
	return len(buf), nil

}

type CryptoReader struct {
	r      io.Reader
	cipher *rc4.Cipher
}

func NewCrytoReader(r io.Reader, key string) io.Reader {
	md5Sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5Sum[:])
	if err != nil {
		log.Print(err)
	}

	return &CryptoReader{
		r:      r,
		cipher: cipher,
	}

}

func (r *CryptoReader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		log.Print(err)
	}
	buf := b[:n]
	r.cipher.XORKeyStream(buf, buf)
	return n, err
}
