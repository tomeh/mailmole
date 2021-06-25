package smtp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/mail"
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
func (server *Server) ListenAndServe(started chan bool, stop chan bool) {
	var stopped bool
	var err error
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Addr, server.Port))
	if err != nil {
		panic(err)
	}

	server.boundPort = listener.Addr().(*net.TCPAddr).Port

	log.Printf("MailMole SMTP Server listening on %s - Port %d\n", server.Addr, server.boundPort)

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

				scanner := bufio.NewScanner(netConn)
				scanner.Split(ScanCRLF)

				smtpConn := connection{
					netConn,
					scanner,
					server,
					session{
						StateInit,
						"",
						"",
						&mail.Message{
							Header: nil,
							Body:   nil,
						},
					},
				}

				smtpConn.handle()
			}()
		}
	}()

	// Notify that the server is started
	started <- true

	// Wait for the stop signal.
	stopped = <-stop

	err = listener.Close()

	// Signal that the server is stopped.
	stop <- true
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
