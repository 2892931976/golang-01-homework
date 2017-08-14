package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"github.com/51reboot/golang-01-homework/lesson10/jcui/mycrypto"
)

//当前函数复制接收数据并进行解密操作
func handleConn_accept(conn net.Conn) {
	defer conn.Close()
	//将数据解密
	key := "AB234asfds345safdasd"
	remote, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Print(err)
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		r := mycrypto.NewCryptoReader(conn, key)
		io.Copy(remote, r)
	}()
	go func() {
		defer wg.Done()
		w := mycrypto.NewCryptoWriter(conn, key)
		io.Copy(w, remote)
	}()
	wg.Wait()
}

//
func mustReadByte(r *bufio.Reader) byte {
	b, err := r.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

//
func handshake(r *bufio.Reader, conn net.Conn) (err error) {
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
	conn.Write(resp)
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
	log.Printf("%s", readtype)
	if readtype != 3 {
		return "", errors.New("bad type")
	}

	addrlen := mustReadByte(r)
	add := make([]byte, addrlen)
	io.ReadFull(r, add)
	log.Printf("addr:%s", add)

	//处理第六个数据,占位2个字节
	var port int16
	binary.Read(r, binary.BigEndian, &port)
	return fmt.Sprintf("%s:%d", add, port), nil
}

//负责处理解密后的数据
func handleConn_real(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	handshake(r, conn)
	addr, err := readAddr(r)
	if err != nil {
		log.Print("Addr Error:", err)
		return
	}
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)
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
		io.Copy(remote, conn)
		remote.Close()
	}()
	//go 接收remote的数据,发送到客户端,直到remote的EOF
	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
		conn.Close()
	}()
	//等待两个协程结束
	wg.Wait()

}

func main() {
	listen_accept, err1 := net.Listen("tcp", ":9999")
	listen_real, err2 := net.Listen("tcp", ":8888")
	if err1 != nil {
		log.Fatal(err1)
	}
	if err2 != nil {
		log.Fatal(err2)
	}
	defer listen_accept.Close()
	defer listen_real.Close()

	go func() {
		for {
			conn_s, err1 := listen_accept.Accept()
			if err1 != nil {
				log.Print(err1)
			}
			go handleConn_accept(conn_s)
			log.Print("s:", conn_s.RemoteAddr().String())
		}
	}()

	func() {
		for {
			conn_d, err2 := listen_real.Accept()
			if err2 != nil {
				log.Print(err2)
			}
			go handleConn_real(conn_d)
			log.Print("d:", conn_d.RemoteAddr().String())

		}
	}()
}
