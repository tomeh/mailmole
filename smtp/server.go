package smtp

import (
	"fmt"
	"log"
	"net"
)

// TODO - DOCUMENT
type Server struct {
	ServerConfig
}

// Create a new Server instance with the given configuration.
// Call method ListenAndServe
// TODO - DOCUMENT
func NewServer(c ServerConfig) Server {
	err := c.validate()
	if err != nil {
		log.Fatalf("ServerConfig invalid - %s", err)
	}
	return Server{c}
}

// TODO - DOCUMENT
func (server *Server) ListenAndServe() error {
	var err error
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Addr, server.Port))
	if err != nil {
		return err
	}

	log.Printf("Mail mole SMTP Server listening on %s - Port %d\n", server.Addr, server.Port)

	for {
		// Wait for new connections.
		netConn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s\n", err)
			continue
		}

		// Dispatch a handler.
		go func() {
			defer func() {
				err := netConn.Close()
				if err != nil {
					log.Printf("Error closing connection %s\n", err)
				}
			}()

			smtpConn := connection{
				netConn,
				server,
				session{StateInit, "", ""},
			}

			smtpConn.handle()
		}()
	}
}

var (
	LogDebug   bool
	version    string
	serverName string
)

func init() {
	LogDebug = true
	version = "0.0.1"
	serverName = "Mailmole"
}
