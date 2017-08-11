package main

import (
	"bufio"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var (
	dir = flag.String("d", "", "file server dir")
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := bufio.NewReader(conn)
	r, err := buf.ReadString('\n')
	if err == io.EOF {
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
	cmd := strings.Fields(strings.TrimSpace(r))
	switch cmd[0] {
	case "GET":
		f, err := os.Open(filepath.Join(*dir, cmd[1]))
		if err != nil {
			log.Println(err)
			conn.Write([]byte(err.Error() + "\n"))
			return
		}
		defer f.Close()
		io.Copy(conn, f)
	case "STORE":
		f, err := os.Create(filepath.Join(*dir, cmd[1]))
		if err != nil {
			log.Println(err)
			conn.Write([]byte(err.Error() + "\n"))
			return
		}
		defer f.Close()
		io.Copy(f, buf)
	case "LS":
		var d string
		if len(cmd) == 1 {
			d = *dir
		} else {
			d = cmd[1]
		}
		info, err := ioutil.ReadDir(d)
		if err != nil {
			log.Println(err)
			conn.Write([]byte(err.Error() + "\n"))
			return
		}
		for _, i := range info {
			conn.Write([]byte(i.Name() + "\n"))
		}
	}
}

func main() {
	flag.Parse()
	addr := ":9000"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConn(conn)
	}

}
