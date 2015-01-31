package main

func main() {

	master := NewMaster()
	slave := NewSlave()
	commandChannel := master.start()
	slave.start(commandChannel)
	//slave(commandChannel)
}

func slave(commandChannel <-chan Command) {
	for {
		cc := <-commandChannel
		cc.exec()
	}
}
