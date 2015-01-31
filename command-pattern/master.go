package main

import (
	"fmt"
)

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
		c := NewPrintHello()
		command <- c

	}()

	return command

}
