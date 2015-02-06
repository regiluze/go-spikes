package main

import (
	"flag"
)

func main() {
	port := flag.String("port", "8080", "listen port")
	server := NewServer(*port)
	server.Start()
}
