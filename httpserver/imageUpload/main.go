package main

import (
	"flag"
	"fmt"

	"github.com/regiluze/go-spikes/httpserver"
)

func main() {
	port := flag.String("port", "8080", "listen port")
	address := flag.String("address", "0.0.0.0", "server address")
	flag.Parse()
	uploadHandler := httpserver.NewImageUploadHandler()
	server := httpserver.NewHttpServer(uploadHandler, *address, *port)
	server.SetErrTemplate(ErrorTemplate)
	error := server.Start()
	if error != nil {
		fmt.Println(error)
	}

}
