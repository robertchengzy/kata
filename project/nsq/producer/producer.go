package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

var producer *nsq.Producer

func main() {
	if len(os.Args) < 3 {
		fmt.Println("参数不足")
		return
	}
	env := os.Args[1]
	param := os.Args[2]

	fmt.Printf("data [%v] [%v]", env, param)

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

	for i := 0; i < 1; i++ {
		data := `{"flexibleTaskRenew":` + param + `}`
		err = producer.Publish("j_pp_phoneBill", []byte(data))
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("send success", i)
	}
}
