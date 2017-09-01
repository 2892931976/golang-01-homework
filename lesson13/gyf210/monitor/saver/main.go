package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/bsm/sarama-cluster"
	"github.com/gyf210/monitor/common"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"time"
)

var (
	kaddr = flag.String("kaddr", "192.168.3.50:9092", "kafka address")
)

func esIndexName(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("20060102")
}

func main() {
	flag.Parse()
	consumer, err := cluster.NewConsumer([]string{*kaddr},
		"falcon-saver", []string{"falcon"}, cluster.NewConfig())
	if err != nil {
		log.Fatalln(err)
	}

	esclient, err := elastic.NewClient(elastic.SetURL("http://192.168.3.50:9200"),
		elastic.SetHealthcheckInterval(10*time.Second))
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case msg := <-consumer.Messages():
			var metric *common.Metric
			err := json.Unmarshal(msg.Value, metric)
			if err != nil {
				log.Println(err)
				continue
			}
			index := esIndexName(metric.Timestamp)
			_, err = esclient.Index().
				Index(index).
				Type("falcon").
				BodyString(string(msg.Value)).
				Do(context.TODO())
			if err != nil {
				log.Println(err)
			}
		case err := <-consumer.Errors():
			log.Println(err)

		}
	}

}
