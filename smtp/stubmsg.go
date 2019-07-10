package smtp

import (
	"fmt"
	"net/mail"
	"strings"
)

func stubMessage(body string) *mail.Message {

	raw := strings.NewReader(fmt.Sprintf(`Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon

%s`, body))

	msg, _ := mail.ReadMessage(raw)
	return msg
}
