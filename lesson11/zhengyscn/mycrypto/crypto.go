package mycrypto

import (
	"crypto/md5"
	"crypto/rc4"
	"io"
	"os"
)

type CryptoWriter struct {
	w      io.Writer
	cipher *rc4.Cipher
}

type CryptoReader struct {
	r      io.Reader
	cipher *rc4.Cipher
}

func NewCryptoWriter(w io.Writer, key string) io.Writer {
	// 接口是方法的集合
	// 实现了这个接口的方法 你就实现了这个接口

	// 接口中有方法签名 只要你的结构体实现了这个方法 那么你就是这个接口
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

// 把b里面的数据进行加密，之后写入到w.w里面
// 调用w.w.Write进行写入
// 过滤器
func (w *CryptoWriter) Write(b []byte) (int, error) {
	// b 只读 默认不可原地修改
	buf := make([]byte, len(b))
	w.cipher.XORKeyStream(buf, b)
	return w.w.Write(buf)
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
	// 尽量把b填满 buf
	n, err := r.r.Read(b)
	buf := b[:n]
	r.cipher.XORKeyStream(buf, buf)
	return n, err
}

func main() {
	// echo "hello" |  ./crypto | ./crypto
	// time ./crypto < block > block.1
	// time ./crypto < block.1 > block.2
	// md5sum
	// 流式处理
	r := NewCryptoReader(os.Stdin, "123456")
	io.Copy(os.Stdout, r)

	//w := NewCryptoWriter(os.Stdout, "123456")
	//io.Copy(w, os.Stdin)

}
