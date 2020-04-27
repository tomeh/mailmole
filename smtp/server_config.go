package smtp

import (
	"fmt"
	"net"
)

// TODO - DOCUMENT
type ServerConfig struct {
	Addr     string
	Port     int
	HostName string
}

func (c *ServerConfig) validate() error {
	if net.ParseIP(c.Addr) == nil {
		return fmt.Errorf("'%s' is not a valid IP address", c.Addr)
	}

	return nil
}
