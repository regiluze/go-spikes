package httpserver

import (
	"fmt"
	"net/http"

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
	port    string
	address string
	handler routeHandler
}

func NewHttpServer(h routeHandler, a string, p string) *HttpServer {
	s := &HttpServer{handler: h, address: a, port: p}
	return s
}

type routeHandler interface {
	HandleRoutes(ErrHandler) *mux.Router
}

type ErrHandler func(http.HandlerFunc) http.HandlerFunc

func (s *HttpServer) errorHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recoverErr := recover(); recoverErr != nil {
				error := NewError(fmt.Sprintf("\"%v\"", recoverErr))
				w.WriteHeader(500)
				errorTemplate.Execute(w, error)
			}
		}()
		fn(w, r)
	}
}

func (s *HttpServer) Start() error {
	r := s.handler.HandleRoutes(s.errorHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/", r)
	return http.ListenAndServe(fmt.Sprintf("%s:%s", s.address, s.port), nil)
}
