package smtp

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

var testServerConfig = ServerConfig{
	"127.0.0.1",
	0,
	"localhost",
}

func TestNewServer(t *testing.T) {
	NewServer(testServerConfig)
}

func TestServer_ListenAndServe(t *testing.T) {
	server := newTestServer()
	server.start()
	defer server.stop()

	client := newTestClient(server)

	if client.greeting != fmt.Sprintf("220 %s ESMTP %s %s ready\r\n", server.HostName, serverName, version) {
		t.Errorf("Greeting is unexpected - %s", client.greeting)
	}
}

func TestSmtp(t *testing.T) {
	server := newTestServer()
	server.start()
	defer server.stop()

	client := newTestClient(server)

	client.Write("HELO example.org")
	heloResponse := client.Read()

	if heloResponse != fmt.Sprintf("250 localhost greets example.org\r\n") {
		t.Errorf("HELO Response is unexpected - %s", heloResponse)
	}

	client.Write("MAIL FROM:<me@example.org>")
	mailFromResponse := client.Read()

	if mailFromResponse != fmt.Sprintf("250 OK\r\n") {
		t.Errorf("MAIL FROM Response is unexpected - %s", mailFromResponse)
	}

	client.Write("RCPT TO:<someone.else@example.org>")
	rcptToResponse := client.Read()

	if rcptToResponse != fmt.Sprintf("250 OK\r\n") {
		t.Errorf("RCPT TO Response is unexpected - %s", rcptToResponse)
	}

	client.Write("DATA\r\n")
	dataResponse := client.Read()
	if dataResponse != fmt.Sprintf("354 Start mail input; end with <CRLF>.<CRLF>\r\n") {
		t.Errorf("DATA Response is unexpected - %s", dataResponse)
	}

	client.Write("Hello SMTP\r\n")
	client.Write(".\r\n")

	endDataResponse := client.Read()

	if endDataResponse != fmt.Sprintf("250 OK\r\n") {
		t.Errorf("DATA (end) Response is unexpected - %s", endDataResponse)
	}
}

func newTestServer() *testServer {
	return &testServer{
		NewServer(testServerConfig),
		make(chan bool),
		make(chan bool),
	}
}

type testServer struct {
	Server
	startCh chan bool
	stopCh chan bool
}

func (ts *testServer) start() {
	go func() {
		err := ts.ListenAndServe(ts.startCh, ts.stopCh)
		if err != nil {
			panic(err)
		}
	}()

	ts.startCh <- true
	<-ts.startCh
}

func (ts *testServer) stop() {
	ts.stopCh <- true
	<-ts.stopCh
}

type testClient struct {
	conn net.Conn
	greeting string
}

func (tc *testClient) Write(msg string) {
	_, _ = tc.conn.Write(append([]byte(msg), []byte("\r\n")...))
}

func (tc *testClient) Read() string {
	response, _ := bufio.NewReader(tc.conn).ReadString('\n')
	return response
}

func newTestClient(server *testServer) *testClient {
	netClient, _ := net.Dial("tcp", fmt.Sprintf("%s:%d", server.ServerConfig.Addr, server.boundPort))
	greeting, _ := bufio.NewReader(netClient).ReadString('\n')

	return &testClient{
		netClient,
		greeting,
	}
}

