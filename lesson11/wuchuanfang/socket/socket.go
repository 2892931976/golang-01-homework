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
)

//1、握手；2、获取客户端的代理请求；3、开始代理
func handshake(r *bufio.Reader, conn net.Conn) error {
	version, _ := r.ReadByte()
	//log.Printf("version:%d", version)
	if version != 5 {
		return errors.New("bad version")
	}
	nmethods, _ := r.ReadByte()
	//log.Printf("nmethods:%d", nmethods)

	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	//log.Printf("%v", buf)

	resp := []byte{5, 0}
	conn.Write(resp)

	return nil
}

func readAddr(r *bufio.Reader) (string, error) {
	version, _ := r.ReadByte()
	if version != 5 {
		return "", errors.New("bad version")
	}
	//b, _ = r.ReadBytes(\n)
	cmd, _ := r.ReadByte()

	if cmd != 1 {
		return "", errors.New("bad cmd")
	}

	r.ReadByte()
	atyp, _ := r.ReadByte()
	//log.Printf("%v%v%v", cmd, rsv, atyp)
	if atyp != 3 {
		return "", errors.New("bad atyp")
	}

	addrlen, _ := r.ReadByte()
	//log.Printf("addrlen:%v", addrlen)
	addr := make([]byte, addrlen)
	io.ReadFull(r, addr)
	//log.Printf("addr:%s", addr)

	var port int16
	binary.Read(r, binary.BigEndian, &port)

	//var host, port string
	/*
	    switch atyp {
	   case 0x01: //IP V4
	       host = net.IPv4(b[4], b[5], b[6], b[7]).String()
	   case 0x03: //域名 host = string(b[5:n-2])//b[4]表示域名的长度
	   case 0x04: //IP V6 host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()
	   }
	   port = strconv.Itoa(int(b[n-2])<<8 | int(b[n-1]))
	*/

	return fmt.Sprintf("%s:%d", addr, port), nil
}

func handleConn(conn net.Conn) {

	//defer conn.Close()
	r := bufio.NewReader(conn)
	handshake(r, conn)
	addr, _ := readAddr(r)
	log.Printf("realaddr:%s", addr)
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} //详情请参考：http://www.jianshu.com/p/172810a70fad
	conn.Write(resp)
	var wg sync.WaitGroup
	wg.Add(2)

	remote, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		conn.Close()
		//remote.Close()
		return
	}
	go func() {
		defer wg.Done()
		io.Copy(remote, conn)
		remote.Close()
	}()
	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
		conn.Close()
	}()
	wg.Wait()
}

func main() {
	//建立监听
	addr := "127.0.0.1:9090"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	//defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConn(conn)
	}

}
