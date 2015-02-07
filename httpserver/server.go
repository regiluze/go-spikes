// Copyright 2015 The httpserver Authors. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.
package httpserver

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type ServerError struct {
	Msg string
}

func NewError(msg string) *ServerError {
	s := &ServerError{Msg: msg}
	return s
}

type HttpServer struct {
	port             string
	address          string
	handler          routeHandler
	errTemplate      *template.Template
	notFoundTemplate *template.Template
}

func NewHttpServer(h routeHandler, a string, p string) *HttpServer {
	s := &HttpServer{handler: h, address: a, port: p}
	return s
}

type routeHandler interface {
	HandleRoutes(ErrHandler) *mux.Router
}

type ErrHandler func(http.HandlerFunc) http.HandlerFunc

func (s *HttpServer) SetErrTemplate(t *template.Template) {
	s.errTemplate = t
}

func (s *HttpServer) SetNotFoundTemplate(t *template.Template) {
	s.notFoundTemplate = t
}

func (s *HttpServer) errorHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recoverErr := recover(); recoverErr != nil {
				fmt.Println("egi code>> recover")
				error := NewError(fmt.Sprintf("\"%v\"", recoverErr))
				w.WriteHeader(500)
				s.errTemplate.Execute(w, error)
			}
		}()
		fn(w, r)
	}
}
func (s *HttpServer) NotFound(w http.ResponseWriter, r *http.Request) {

	fmt.Println("egi not found")
	w.WriteHeader(404)
	s.notFoundTemplate.Execute(w, nil)

}

func (s *HttpServer) Start() error {
	r := s.handler.HandleRoutes(s.errorHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/", r)
	r.NotFoundHandler = http.HandlerFunc(s.NotFound)
	return http.ListenAndServe(fmt.Sprintf("%s:%s", s.address, s.port), nil)
}
