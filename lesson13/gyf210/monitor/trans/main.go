package main

import (
	"bufio"
	"flag"
	"github.com/Shopify/sarama"
	"log"
	"net"
)

var (
	addr  = flag.String("addr", ":6000", "listen address")
	kaddr = flag.String("kaddr", "192.168.3.50:9092", "kafka address")
)

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	producer, err := sarama.NewAsyncProducer([]string{*kaddr}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	ch := producer.Input()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn, ch)
	}
}

func handle(conn net.Conn, ch chan<- *sarama.ProducerMessage) {
	defer conn.Close()

	buf := bufio.NewReader(conn)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1]
		message := &sarama.ProducerMessage{
			Topic: "falcon",
			Key:   nil,
			Value: sarama.StringEncoder(line),
		}
		ch <- message
	}
}
