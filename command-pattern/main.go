package main

import (
	"fmt"
	"time"
)

type Command interface {
	exec()
}

type PrintHello struct {
}

func (p *PrintHello) exec() {
	for {
		fmt.Println("kaixo")
		time.Sleep(5 * time.Second)
		fmt.Println("agur")
	}
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
	for {
		cc := <-commandChannel
		cc.exec()
	}
}
