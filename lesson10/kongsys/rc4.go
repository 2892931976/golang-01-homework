package main

import (
	"crypto/rc4"
	"io"
	"log"
)

func crypto(w io.Writer, r io.Reader, key string) {
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
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
	key := "123456"
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	buf := []byte("hello")
	cipher.XORKeyStream(buf, buf)
	log.Printf(string(buf))
}
