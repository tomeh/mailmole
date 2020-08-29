package smtp

import (
	"os"
	"strings"
	"testing"
)

var addr = "0.0.0.0"
var port = 0
var hostname string

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		// TODO
		panic("err")
	}
}

// Ensure that a validate ServerConfig is validated.
func TestServerConfigValidate(t *testing.T) {
	serverConfig := &ServerConfig{
		addr,
		port,
		hostname,
	}

	err := serverConfig.validate()

	if err != nil {
		t.Errorf("Expected ServerConfig to validate OK but it returned an error - %s", err)
	}
}

// Ensure that invalid IP addresses are caught.
func TestServerConfigValidateFails(t *testing.T) {
	serverConfig := &ServerConfig{
		"one.two.thr.fou",
		port,
		hostname,
	}

	err := serverConfig.validate()

	if err == nil || !strings.Contains(err.Error(), "addr") {
		t.Error("ServerConfig should have returned an 'addr' error")
	}
}
