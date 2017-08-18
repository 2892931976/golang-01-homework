package mycrypto

import (
	"crypto/md5"
	"crypto/rc4"
	"io"
	"log"
	"os"
)

// CryptoWriter crypto Writer
type CryptoWriter struct {
	w      io.Writer
	cipher *rc4.Cipher
}

// NewCryptoWriter new Writer crypto by key
// Writer是实现了Write方法的Interface接口。返回的结构体实现了Write方法，
// 因此，还是一个Writer
func NewCryptoWriter(w io.Writer, key string) io.Writer {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher([]byte(md5sum[:]))
	if err != nil {
		log.Panic(err)
	}
	return &CryptoWriter{
		w:      w,
		cipher: cipher,
	}
}

// CryptoReader crypto Reader
type CryptoReader struct {
	r      io.Reader
	cipher *rc4.Cipher
}

// NewCryptoReader new Reader decrypto by key
func NewCryptoReader(r io.Reader, key string) io.Reader {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher([]byte(md5sum[:]))
	if err != nil {
		log.Panic(err)
	}
	return &CryptoReader{
		r:      r,
		cipher: cipher,
	}
}

// Write for CryptoWriter
// Write writes len(b) bytes from b to the underlying data stream
// 把b里面的数据进行加密，之后写入到w.w（调用w.w.Write()
func (w *CryptoWriter) Write(b []byte) (int, error) {
	// b是数据源，通常建议自定义一个buf，从而不修改b。
	bLen := len(b)
	buf := make([]byte, bLen)
	w.cipher.XORKeyStream(buf, b[:bLen])
	return w.w.Write(buf)
}

// Read for CryptoReader
// Read reads up to len(b) bytes into b
func (r *CryptoReader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	buf := b[:n]

	r.cipher.XORKeyStream(buf, buf)
	return n, err
}

func main() {
	var key = "123456"
	r := NewCryptoReader(os.Stdin, key)
	io.Copy(os.Stdout, r)

	w := NewCryptoWriter(os.Stdout, key)
	io.Copy(w, os.Stdin)
}
