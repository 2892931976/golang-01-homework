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
	&cipher.StreamReaderu
}
