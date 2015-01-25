package main

import "fmt"

type Command interface {
	exec()
}

type PrintHello struct {
}

func (p *PrintHello) exec() {

	fmt.Println("kaixo")

}

func NewPrintHello() Command {
	return &PrintHello{}
}

func main() {

	commandChannel := startCommander()
	slave(commandChannel)
}

func startCommander() <-chan Command {

	fmt.Println("commander started")
	command := make(chan Command)

	go func() {
		c := NewPrintHello()
		command <- c

	}()

	return command

}

func slave(commandChannel <-chan Command) {
	cc := <-commandChannel
	cc.exec()

}
