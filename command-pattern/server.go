package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	s := &Server{}
	return s

}

type Data struct {
	Msg string
}

type CommandHandler struct {
	CommandChannel chan Command
}

func NewCommandHandler() *CommandHandler {

	ch := &CommandHandler{CommandChannel: make(chan Command)}
	return ch

}

func (ch *CommandHandler) handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		decoder := json.NewDecoder(r.Body)
		var d Data
		err := decoder.Decode(&d)
		if err != nil {
			panic(err)
		}
		fmt.Println(d)
		c := NewPrintHello(d.Msg)
		ch.CommandChannel <- c
	}
}

func (s *Server) start() <-chan Command {
	ch := NewCommandHandler()
	go func() {
		http.HandleFunc("/commander", ch.handler)
		http.ListenAndServe(":8080", nil)
	}()
	return ch.CommandChannel
}
