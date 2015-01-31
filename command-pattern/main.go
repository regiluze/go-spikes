package main

func main() {
	slave := NewSlave()
	server := NewServer()
	commandChannel := server.start()
	slave.start(commandChannel)
}
