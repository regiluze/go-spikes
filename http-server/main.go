package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.String("port", "8080", "listen port")
	address := flag.String("address", "0.0.0.0", "server address")
	flag.Parse()
	uploadHandler := NewImageUploadHandler()
	server := NewServer(uploadHandler, *address, *port)
	error := server.Start()
	if error != nil {
		fmt.Println(error)
	}

}
