package smtp

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
)

// Represents a connection from an SMTP Client.
type connection struct {
	rw net.Conn
	*bufio.Scanner
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

	for c.Scan() {
		if handleClientInput(c.Text(), c) {
			break
		}
	}
}

func ScanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte{'\r','\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0:i]), nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}

	// Request more data.
	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}

	return data
}
