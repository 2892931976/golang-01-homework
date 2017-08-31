package main

import (
	"log"

	cluster "github.com/bsm/sarama-cluster"
	elastic "gopkg.in/olivere/elastic.v5"
)

func main() {
	consumer, err := cluster.NewConsumer(
		[]string{":9092"},
		"falcon-saver",
		[]string{"falcon"},
		cluster.NewConfig(),
	)

	if err != nil {
		log.Fatal(err)
	}

	esclient, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case msg := <-consumer.Messages():
			_, err esclient.Index().
				Type(indexName()).
				BodyString(string(msg.Value)).
				Do(context.TODO())
			if err!= nil{
				log.Print(err)
			}
			log.Print(string(msg.Value))
		case err := <-consumer.Errors():
			log.Print(err)
		}
	}
}
