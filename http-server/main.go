package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"text/template"
)

var uploadTemplate = template.Must(template.ParseFiles("index.html"))

func handle(w http.ResponseWriter, r *http.Request) {
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

func main() {
	fmt.Println("kaixo")
	http.HandleFunc("/", handle)
	http.HandleFunc("/view", view)
	http.ListenAndServe(":8080", nil)
	fmt.Println("agur")

}
