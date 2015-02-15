package main

import (
	"encoding/json"
	//	"fmt"
	"time"

	"github.com/aleasoluciones/simpleamqp"
)

type CommandMessage struct {
	Msg string `json:"msg"`
}

type AmqpMaster struct {
	consumer *simpleamqp.AmqpConsumer
}

func NewAmqpMaster(amqpuri string) *AmqpMaster {
	amqpConsumer := simpleamqp.NewAmqpConsumer(amqpuri)
	a := &AmqpMaster{consumer: amqpConsumer}
	return a
}

func (a *AmqpMaster) start(cf CommandFactory) <-chan Command {
	ch := NewCommandHandler(cf)

	go func() {
		commands := a.consumer.Receive("events",
			[]string{"#", "test"},
			"", simpleamqp.QueueOptions{Durable: false, Delete: true, Exclusive: true},
			30*time.Minute)
		for command := range commands {
			var m CommandMessage
			err := json.Unmarshal([]byte(command.Body), &m)
			if err == nil {
				c := NewPrintHello(m.Msg)
				ch.CommandChannel <- c
			}
		}
	}()

	return ch.CommandChannel
}
