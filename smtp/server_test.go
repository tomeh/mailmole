package smtp

import "testing"

var config = ServerConfig{
	"127.0.0.1",
	25,
}

func TestNewServer(t *testing.T) {
	NewServer(config)
}

func TestServer_ListenAndServe(t *testing.T) {
	// TODO - TEST
}
