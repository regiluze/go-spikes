package main

import (
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	s := &Server{}
	return s

}

func (s *Server) start(cf CommandFactory) <-chan Command {
	ch := NewCommandHandler(cf)
	go func() {
		http.HandleFunc("/commander", ch.handler)
		http.ListenAndServe(":8080", nil)
	}()
	return ch.CommandChannel
}
