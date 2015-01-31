package main

import (
	"fmt"
	"time"
)

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
