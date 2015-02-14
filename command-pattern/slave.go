package main

type Slave struct {
	mainChan chan Command
}

func NewSlave() *Slave {
	s := &Slave{}
	return s
}

func (s *Slave) AddMaster(cs ...<-chan Command) {
	s.mainChan = make(chan Command)
	for _, c := range cs {
		s.mainChan <- c
	}
}

func (s *Slave) Start() {
	for {
		cc := <-s.mainChan
		cc.exec()
	}
}
