package main

import (
//	"fmt"
)

type Slave struct {
	mainChan chan Command
}

func NewSlave() *Slave {
	s := &Slave{}
	return s
}

func (s *Slave) bindMainChan(c <-chan Command) {
	for cmd := range c {
		s.mainChan <- cmd
	}
}

func (s *Slave) AddMaster(cs ...<-chan Command) {
	s.mainChan = make(chan Command)
	for _, c := range cs {
		go s.bindMainChan(c)
	}
}

func (s *Slave) Start() {
	for {
		cc := <-s.mainChan
		cc.exec()
	}
}
