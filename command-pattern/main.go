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

func main() {

	master := NewMaster()
	commandChannel := master.start()
	slave(commandChannel)
}

func slave(commandChannel <-chan Command) {
	for {
		cc := <-commandChannel
		cc.exec()
	}
}
