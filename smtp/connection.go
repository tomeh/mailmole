package smtp

import (
	"fmt"
	"net"
	"strings"
)

var lfReplacer *strings.Replacer

func init() {
	lfReplacer = strings.NewReplacer("\r\n", "\r\n", "\r", "\r\n")
}

// Represents a connection from an SMTP Client.
type connection struct {
	rw net.Conn
	*Server
	session
}

// Send the message to the connection. If the message starts with '-'
// then this takes the place of the space between the status and the message.
func (c *connection) send(status int, message string) error {
	spacer := " "
	if message[:1] == "-" {
		spacer = "-"
		message = message[1:]
	}
	_, err := c.rw.Write([]byte(fmt.Sprintf("%d%s%s\r\n", status, spacer, message)))
	return err
}

// Send the message to the connection, or panic on error
// See c.Send()
func (c *connection) mustSend(status int, message string) {
	if err := c.send(status, message); err != nil {
		panic(err)
	}
}

// Handle the connection. First a greeting is sent, then lines read in from the underlying
// TCP connection. Once we have a line of text, delimited by <CRLF>, we pass it to a handler.
func (c *connection) handle() {
	c.mustSend(220, fmt.Sprintf("%s ESMTP %s %s ready", c.Server.HostName, serverName, version))

	buf := make([]byte, 4)
	var current string
	for {
		// Read the next bytes into the buffer.
		n, err := c.rw.Read(buf)
		if err != nil {
			panic(err)
		}

		// Standardise any LF chars with <CRLF>.
		if len(buf[:n]) >= 2 && string(buf[n-2:n]) == "\r\n" {
			// If we already have an <CRLF> at the end then leave it.
			current += lfReplacer.Replace(string(buf[:n-2]))
			current += "\r\n"
		} else {
			// Don't include the final char as we could potentially be cutting a <CRLF>
			// halfway through. We add the final character back in on the following statement.
			current += lfReplacer.Replace(string(buf[:n-1]))
			current += string(buf[n-1:n])
		}

		// Find the first <CRLF> if present.
		// There will only be 1 part if there is no <CRLF> in the current line.
		parts := strings.SplitN(current, "\r\n", 2)
		if len(parts) > 1 {
			// Handle the line (first part) and assign the remainder of the line to "current".
			handleClientInput(string(parts[0]), c)
			current = parts[1]
		}

		if c.session.State == StateQuit {
			break
		}
	}
}
