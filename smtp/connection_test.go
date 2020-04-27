package smtp

import (
	"fmt"
	"net/mail"
	"strings"
	"testing"
)

func TestConnection_handle(t *testing.T) {

}

func stubMessage(subject, body string) *mail.Message {
	raw := strings.NewReader(fmt.Sprintf(`Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: %s

%s`, subject, body))

	msg, _ := mail.ReadMessage(raw)
	return msg
}
