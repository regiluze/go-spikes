package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CommandFactory interface {
	Get(string) Command
}

type Data struct {
	Msg string
}

type CommandHandler struct {
	CommandChannel chan Command
	Factory        CommandFactory
}

func NewCommandHandler(cf CommandFactory) *CommandHandler {

	ch := &CommandHandler{CommandChannel: make(chan Command),
		Factory: cf}
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
		c := ch.Factory.Get(d.Msg)
		ch.CommandChannel <- c
	}
}

func (ch *CommandHandler) HandleRoutes(errFunc errHandler) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", errFunc(ch.handler))
	return r

}
