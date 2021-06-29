package smtp

import (
	"time"

	"github.com/tomedharris/mailmole/contracts"
)

func NewServer(sub <-chan contracts.Listener) *Server {
	s := &Server{}

	go func(sub <-chan contracts.Listener) {
		for incoming := range sub {
			s.subscribe(incoming)
		}
	}(sub)

	return s
}

type Server struct {
	listeners []contracts.Listener
}

func (s *Server) Start() {
	for {
		for _, l := range s.listeners {
			l.Publish("Hi")
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func (s *Server) subscribe(l contracts.Listener) {
	s.listeners = append(s.listeners, l)
}
