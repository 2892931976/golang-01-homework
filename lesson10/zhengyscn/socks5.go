package main

import (
	"bufio"
	"errors"

	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

func readAddr(r *bufio.Reader) (string, error) {
	version, err := r.ReadByte()
	log.Printf("version:%d", version)
	if err != nil {
		return "", err
	}
	if version != 5 {
		return "", errors.New("bad version")
	}

	cmd, err := r.ReadByte()
	log.Printf("cmd:%d", cmd)
	if err != nil {
		return "", err
	}

	r.ReadByte()

	addratyp, _ := r.ReadByte()
	log.Printf("addratyp:%s", addratyp)
	if addratyp != 3 {
		return "", errors.New("bad addratyp")
	}

	atyplen, err := r.ReadByte()
	log.Printf("atyplen:%d", atyplen)
	if err != nil {
		return "", err
	}
	buf := make([]byte, atyplen)
	io.ReadFull(r, buf)
	log.Printf("buf:%s", buf)

	var port uint16
	binary.Read(r, binary.BigEndian, &port)

	addr := buf
	return fmt.Sprintf("%s:%d", addr, port), nil
}

func handshake(r *bufio.Reader, conn net.Conn) error {
	version, err := r.ReadByte()
	log.Printf("version:%d", version)
	if err != nil {
		return err
	}
	if version != 5 {
		return errors.New("bad version")
	}
	nmethods, err := r.ReadByte()
	log.Printf("nmethods:%d", nmethods)
	if err != nil {
		return err
	}
	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	log.Printf("buf:%v", buf)

	conn.Write([]byte{5, 0})
	return nil

}

func handConn(conn net.Conn) {
	defer conn.Close()
	if conn == nil {
		return
	}

	r := bufio.NewReader(conn)
	err := handshake(r, conn)
	if err != nil {
		log.Print(err)
		return
	}

	addr, _ := readAddr(r)
	log.Printf("addr:%s", addr)

	serverConn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Print(err)
		return
	}
	defer serverConn.Close()
	conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //响应客户端连接成功

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		io.Copy(serverConn, conn)
		serverConn.Close()
	}()

	go func() {
		defer wg.Done()
		io.Copy(conn, serverConn)
		conn.Close()
	}()
	wg.Wait()

}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		log.Printf("new connect %s\n", conn.LocalAddr().String())
		if err != nil {
			log.Fatal(err)
		}
		go handConn(conn)
	}
}
