package main

import (
	//"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

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
	//f, _, err := r.FormFile("image")
	//check(err)
	f, err := os.Open("filename.ext")
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
				//io.WriteString(w, recoverErr)
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
