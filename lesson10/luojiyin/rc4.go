package main

import (
	"crypto/md5"
	"crypto/rc4"
	"io"
	"log"
)

func crypto(w io.Writer, r io.Reader, key string) {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc.NewCipher([]byte(md5sum))
	buf := make([]byte, 4096)

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}

		cipher.XORKeyStream(buf, buf)
		w.Write(buf)
	}

}

func main() {
	key := "luojiyin"
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	buf := []byte("hello")

	cipher.XORKeyStream(buf, buf)

}
