package main

import (
	"fmt"

	"github.com/regiluze/go-spikes/httpserver"
)

const (
	ADDRESS = "0.0.0.0"
	PORT    = "8080"
)

type CommandServer struct {
}

func NewCommandServer() *CommandServer {
	s := &CommandServer{}
	return s

}

func (s *CommandServer) start(cf CommandFactory) <-chan Command {
	ch := NewCommandHandler(cf)
	go func() {
		httpserver := httpserver.NewHttpServer(ch, ADDRESS, PORT)
		err := httpserver.Start()
		if err != nil {
			fmt.Println(err)
		}
		//http.HandleFunc("/commander", ch.handler)
		//http.ListenAndServe(":8080", nil)
	}()
	return ch.CommandChannel
}
