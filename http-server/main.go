package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "kaixo")
}

func main() {

	http.HandleFunc("/", handle)
	http.ListenAndServe("8080:", nil)

}
