package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Shopify/sarama"
)

// AddrInfo for listen and remote addr
type AddrInfo struct {
	SaverAddr  string `toml:"saver_addr"`
	ListenAddr string `toml:"listen_addr"`
}

var (
	cfgPath  = flag.String("config", "../config/trans.toml", "tranfer config path")
	addrinfo AddrInfo
)

func main() {
	_, err := toml.DecodeFile(*cfgPath, &addrinfo)
	if err != nil {
		log.Fatal("decode toml err:", err)
	}
	log.Printf("addr info:%#v", addrinfo)

	lisn, err := net.Listen("tcp", addrinfo.ListenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer lisn.Close()

	log.Print("saver addr:", addrinfo.SaverAddr)
	producer, err := sarama.NewAsyncProducer([]string{addrinfo.SaverAddr}, nil)
	if err != nil {
		log.Fatal(err)
	}
	ch := producer.Input()

	for {
		conn, err := lisn.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, ch)
	}
}

// 按行读取并处理。ch是只写性质的chan
func handleConn(conn net.Conn, ch chan<- *sarama.ProducerMessage) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err == io.EOF {
			time.Sleep(time.Millisecond * 100)
		}
		if len(line) == 0 {
			continue
		}
		line = line[:len(line)-1]

		msg := &sarama.ProducerMessage{
			Topic: "falcon",
			Key:   nil,
			Value: sarama.StringEncoder(line),
		}
		ch <- msg
		log.Print("kafka msg:", msg)
	}
}
