package main

import (
	"fmt"
)

type PrintHello struct {
	Msg string
}

func (p *PrintHello) exec() {
	fmt.Println("command:", p.Msg)
}

func NewPrintHello(message string) Command {
	return &PrintHello{Msg: message}
}

type Master struct{}

type Command interface {
	exec()
}

func NewMaster() *Master {
	m := &Master{}
	return m
}

func (m *Master) start() <-chan Command {
	fmt.Println("commander started")
	command := make(chan Command)

	go func() {
		c := NewPrintHello("static")
		command <- c

	}()

	return command

}
