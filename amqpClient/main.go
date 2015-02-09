package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/aleasoluciones/simpleamqp"
)

type Message struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func main() {
	amqp := flag.String("uri", "amqp://guest:guest@localhost/", "AMQP connection uri")
	exchange := flag.String("exchange", "events", "Queue exchange")
	topic := flag.String("topic", "test", "topic")
	flag.Parse()

	m := &Message{Name: "", Id: ""}
	json, _ := json.Marshal(m)

	amqpPublisher := simpleamqp.NewAmqpPublisher(*amqp, *exchange)
	amqpPublisher.Publish(*topic, []byte(json))

	fmt.Println("message published!", m)

}
