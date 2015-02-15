package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/regiluze/httpserver"
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

func (ch *CommandHandler) GetRoutes() []*httpserver.Route {
	root := httpserver.NewRoute("", ch.handler)
	routes := []*httpserver.Route{root}
	return routes
}
