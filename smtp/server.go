package smtp

import (
	"fmt"
	"log"
	"net"
)

// TODO - DOCUMENT
type Server struct {
	ServerConfig
	boundPort int // Because port may have been decided by the system.
}

// Create a new Server instance with the given configuration.
// Call method ListenAndServe
// TODO - DOCUMENT
func NewServer(c ServerConfig) Server {
	err := c.validate()
	if err != nil {
		log.Fatalf("ServerConfig invalid - %s", err)
	}
	return Server{c, 0}
}

// TODO - DOCUMENT
func (server *Server) ListenAndServe(start chan bool, stop chan bool) error {
	var stopped bool
	var err error
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Addr, server.Port))
	if err != nil {
		return err
	}

	server.boundPort = listener.Addr().(*net.TCPAddr).Port

	log.Printf("MailMole SMTP Server listening on %s - Port %d\n", server.Addr, server.boundPort)

	// Wait for the start signal.
	<-start
	go func() {
		for {
			// Wait for new connections.
			netConn, err := listener.Accept()

			if stopped {
				return
			}

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
	}()

	// Signal that the server is started.
	start <- true

	// Wait for the stop signal.
	stopped = <-stop

	err = listener.Close()

	// Signal that the server is stopped.
	stop <- true

	return err
}

var (
	LogDebug   bool
	version    string
	serverName string
)

func init() {
	LogDebug = true
	version = "0.0.1"
	serverName = "MailMole"
}
