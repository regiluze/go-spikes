package main

import "fmt"

type Command interface{}

func NewCommand() {
	return Command{}
}

func main() {

	fmt.Println("kaixo")
	command := startCommander()

}

func startCommander() <-chan Command {

	fmt.Println("commander started")
	command := make(chan Command)

	go func() {
		c := NewCommand{}
		command <- c

	}()

	return command

}
