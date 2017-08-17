package main

import (
	//"crypto/md5"
	"crypto/md5"
	"crypto/rc4"
	"flag"
	"io"
	"log"
	"os"
)

var (
	key = flag.String("k", "", "secret key")
)

func crypto(w io.Writer, r io.Reader, key string) {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		cipher.XORKeyStream(buf[:n], buf[:n])
		log.Println(string(buf[:n]))
		w.Write(buf[:n])
	}
}

func main() {
	flag.Parse()
	crypto(os.Stdout, os.Stdin, *key)
}
