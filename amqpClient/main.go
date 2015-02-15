package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/aleasoluciones/simpleamqp"
)

type Message struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Msg  string `json:msg`
}

func main() {
	amqp := flag.String("uri", "amqp://guest:guest@localhost/", "AMQP connection uri")
	exchange := flag.String("exchange", "events", "Queue exchange")
	topic := flag.String("topic", "test", "topic")
	flag.Parse()

	m := &Message{Name: "", Id: "", Msg: "kaxo from rabbit"}
	json, _ := json.Marshal(m)
	fmt.Println("exchage:", *exchange)
	fmt.Println("topic:", *topic)
	amqpPublisher := simpleamqp.NewAmqpPublisher(*amqp, *exchange)
	amqpPublisher.Publish(*topic, []byte(json))

	fmt.Println("message published!", m)

	time.Sleep(2 * time.Second)

}
