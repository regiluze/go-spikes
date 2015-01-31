package main

func main() {

	server := NewServer()
	server.start()
	master := NewMaster()
	slave := NewSlave()
	commandChannel := master.start()
	slave.start(commandChannel)
}

func slave(commandChannel <-chan Command) {
	for {
		cc := <-commandChannel
		cc.exec()
	}
}
