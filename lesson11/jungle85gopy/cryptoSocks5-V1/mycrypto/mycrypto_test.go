package mycrypto

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

// 基准测试
func TestCrypto(t *testing.T) {
	key := "123456"
	testStr := "hello world"

	memBuf := new(bytes.Buffer)
	rdBuf := make([]byte, 1024)

	// 基于buffer构造接口。
	w := NewCryptoWriter(memBuf, key)
	r := NewCryptoReader(memBuf, key)

	// write and read
	w.Write([]byte(testStr))
	n, _ := r.Read(rdBuf)

	if string(rdBuf[:n]) != testStr {
		t.Errorf("not equal:%s, %s", rdBuf[:n], testStr)
	}
}

// 性能测试
func BenchmarkCrypto(b *testing.B) {
	buf := []byte(strings.Repeat("a", 1024))
	key := "123456"

	w := NewCryptoWriter(ioutil.Discard, key)
	// 	按块加密
	for i := 0; i < b.N; i++ {
		n, _ := w.Write(buf)
		// 记录在一个操作中处理的字节数
		b.SetBytes(int64(n))
	}
}
