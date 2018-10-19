package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"strconv"
	"time"
)

var producer *nsq.Producer

func main() {
	nsqd := "192.168.1.51:4150"
	var err error
	producer, err = nsq.NewProducer(nsqd, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	err = producer.Ping()
	if nil != err {
		producer.Stop()
		panic(err)
	}

	for i := 0; i < 100; i++ {
		err = producer.Publish("test", []byte("nihao"+strconv.Itoa(i)))
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("send success", i)
		time.Sleep(3 * time.Second)
	}
}
