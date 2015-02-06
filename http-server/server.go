package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"text/template"
)

type ServerError struct {
	Msg string
}

func NewError(msg string) *ServerError {

	s := &ServerError{Msg: msg}
	return s

}

var uploadTemplate = template.Must(template.ParseFiles("index.html"))
var errorTemplate = template.Must(template.ParseFiles("error.html"))

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		uploadTemplate.Execute(w, nil)
		return
	}
	f, _, err := r.FormFile("image")
	check(err)
	defer f.Close()
	t, err := ioutil.TempFile(".", "image-")
	check(err)
	defer t.Close()
	_, copyErr := io.Copy(t, f)
	check(copyErr)
	http.Redirect(w, r, "/view?id="+t.Name()[6:], 302)
}

func view(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, "image-"+r.FormValue("id"))
}

func errorHandler(fn http.HandlerFunc) http.HandlerFunc {
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

type Server struct {
	port    string
	address string
}

func NewServer(a string, p string) *Server {
	s := &Server{address: a, port: p}
	return s
}

func (s *Server) Start() error {
	fmt.Println("egi", s.address, s.port)
	http.HandleFunc("/", errorHandler(upload))
	http.HandleFunc("/view", errorHandler(view))
	return http.ListenAndServe(fmt.Sprintf("%s:%s", s.address, s.port), nil)
}
