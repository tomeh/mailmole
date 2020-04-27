package smtp

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// TODO - DOCUMENT
type connection struct {
	rw net.Conn
	*Server
	session
}

// TODO - DOCUMENT
func (c *connection) send(status int, message string) error {
	spacer := " "
	if message[:1] == "-" {
		spacer = "-"
		message = message[1:]
	}
	_, err := c.rw.Write([]byte(fmt.Sprintf("%d%s%s\n", status, spacer, message)))
	return err
}

func (c *connection) mustSend(status int, message string) {
	if err := c.send(status, message); err != nil {
		panic(err)
	}
}

// TODO - DOCUMENT
func (c *connection) handle() {
	c.mustSend(220, fmt.Sprintf("%s ESMTP %s %s ready", c.Server.HostName, serverName, version))

	for {
		func() {
			bufReader := bufio.NewReader(c.rw)
			input, _ := bufReader.ReadString('\n')
			if LogDebug {
				log.Printf("Received input %s", input)
			}
			handleClientInput(input, c)
		}()

		if c.session.State == StateQuit {
			break
		}
	}
}
