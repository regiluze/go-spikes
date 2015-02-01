package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

var uploadTemplate = template.Must(template.ParseFiles("index.html"))
var errorTemplate = template.Must(template.ParseFiles("error.html"))

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		uploadTemplate.Execute(w, nil)
		return
	}
	f, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()
	t, err := ioutil.TempFile(".", "image-")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer t.Close()
	if _, err := io.Copy(t, f); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
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
				fmt.Fprintf(os.Stderr, "were panicing: %s\n", recoverErr)
				w.WriteHeader(500)
				errorTemplate.Execute(w, recoverErr)
			}
		}()
		fn(w, r)
	}
}
func main() {
	http.HandleFunc("/", errorHandler(upload))
	http.HandleFunc("/view", errorHandler(view))
	http.ListenAndServe(":8080", nil)
}
