package main

type Slave struct {
}

func NewSlave() *Slave {
	s := &Slave{}
	return s
}

func (s *Slave) start(commandChannel <-chan Command) {
	for {
		cc := <-commandChannel
		cc.exec()
	}
}
