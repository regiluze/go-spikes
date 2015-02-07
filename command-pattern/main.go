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
	httpSlave := NewSlave()
	commandFactory := NewPrintHelloCommandFactory()
	server := NewCommandServer()
	httpCommandChannel := server.start(commandFactory)
	httpSlave.start(httpCommandChannel)
}
