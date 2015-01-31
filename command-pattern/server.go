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

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		decoder := json.NewDecoder(r.Body)
		var d Data
		err := decoder.Decode(&d)
		if err != nil {
			panic(err)
		}
		fmt.Println(d)
		fmt.Println("msg:", d.Msg)
	}
}

func (s *Server) start() {
	go func() {
		http.HandleFunc("/commander", handler)
		http.ListenAndServe(":8080", nil)
	}()

}
