package main

import (
	"fmt"

	"github.com/regiluze/httpserver"
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
		httpserver := httpserver.NewHttpServer(ADDRESS, PORT)
		httpserver.DeployAtBase(ch)
		err := httpserver.Start()
		if err != nil {
			fmt.Println(err)
		}
	}()
	return ch.CommandChannel
}
