package main

type PrintHelloCommandFactory struct {
}

func NewPrintHelloCommandFactory() *PrintHelloCommandFactory {

	commandFactory := &PrintHelloCommandFactory{}
	return commandFactory

}

func (commandFactory *PrintHelloCommandFactory) Get(msg string) Command {

	c := NewPrintHello(msg)
	return c

}

func main() {
	commandFactory := NewPrintHelloCommandFactory()
	server := NewCommandServer()
	httpCommandChannel := server.start(commandFactory)

	amqpMaster := NewAmqpMaster("amqp://guest:guest@localhost/")
	amqpCommandChannel := amqpMaster.start(commandFactory)

	httpSlave := NewSlave()
	httpSlave.AddMaster(httpCommandChannel, amqpCommandChannel)
	httpSlave.Start()
}
