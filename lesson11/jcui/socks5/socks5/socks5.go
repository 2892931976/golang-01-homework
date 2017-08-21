package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/51reboot/golang-01-homework/lesson10/jcui/mycrypto"
	"io"
	"log"
	"net"
	"sync"
)

//
func mustReadByte(r *bufio.Reader) byte {
	b, err := r.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

//
func handshake(r *bufio.Reader, w io.Writer) (err error) {
	defer func() {
		e := recover() // interface{}
		if e != nil {
			err = e.(error)
		}
	}()
	version := mustReadByte(r) //ReadByte 代表读取一个字节
	//处理第一个字节
	log.Printf("version:%d", version)
	if version != 5 {
		return errors.New("bad version")
	}
	//处理第二个字节
	nmethods := mustReadByte(r)
	buf := make([]byte, nmethods)
	io.ReadFull(r, buf) //将buf填充满

	//返回数据
	resp := []byte{5, 0}
	w.Write(resp)
	return nil
}

//
func readAddr(r *bufio.Reader) (addr string, err error) {
	defer func() {
		e := recover() // interface{}
		if e != nil {
			err = e.(error)
		}
	}()
	//处理第一个数据
	version := mustReadByte(r)
	if version != 5 {
		return "", errors.New("bad version")
	}

	//处理第二个数据
	cmd := mustReadByte(r)
	if err != nil {
		return "", err
	}
	if cmd != 1 {
		return "", errors.New("bad cmd")
	}
	//处理第三个数据(保留数据跳过即可)
	mustReadByte(r)

	//处理第四个数据
	readtype := mustReadByte(r)
	//log.Printf("%s", readtype)
	if readtype != 3 {
		return "", errors.New("bad type")
	}

	addrlen := mustReadByte(r)
	addrs := make([]byte, addrlen)
	io.ReadFull(r, addrs)
	//log.Printf("addr:%s", addrs)

	//处理第六个数据,占位2个字节
	var port int16
	binary.Read(r, binary.BigEndian, &port)
	return fmt.Sprintf("%s:%d", addrs, port), nil
}

//负责处理解密后的数据
func handleConn(conn net.Conn) {
	defer conn.Close()
	key := "AB234asfds345safdasd"
	r := bufio.NewReader(mycrypto.NewCryptoReader(conn, key))
	w := mycrypto.NewCryptoWriter(conn, key)
	handshake(r, w)
	addr, err := readAddr(r)
	if err != nil {
		log.Print("Addr Error:", err)
		return
	}
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	w.Write(resp)
	remote, err := net.Dial("tcp", addr)
	if err != nil {
		log.Print("Error:", err)
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	//go 接收客户端的数据,发送到remote,直到conn的EOF
	go func() {
		defer wg.Done()
		io.Copy(remote, r)
		remote.Close()
	}()
	//go 接收remote的数据,发送到客户端,直到remote的EOF
	go func() {
		defer wg.Done()
		io.Copy(w, remote)
		conn.Close()
	}()
	//等待两个协程结束
	wg.Wait()

}

func main() {
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConn(conn)
		log.Print("s:", conn.RemoteAddr().String())
	}
}
