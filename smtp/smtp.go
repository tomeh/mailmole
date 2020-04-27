package smtp

import (
	"fmt"
	"regexp"
	"strings"
)

type State string

type session struct {
	State
}

func (s *session) changeState(state State) {
	s.State = state
}

type handler struct {
	match
	clientInputHandler
}

type match func(string) bool

type clientInputHandler func(string, *connection)

var handlers []*handler

func init() {
	handlers = append(
		handlers,
		&handler{
			func(input string) bool {
				return len(input) >= 4 && strings.ToUpper(input[:4]) == "HELO"
			},
			onHelo,
		},
		&handler{
			func(input string) bool {
				return len(input) >= 4 && strings.ToUpper(input[:4]) == "EHLO"
			},
			onEhlo,
		},
		&handler{
			func(input string) bool {
				return len(input) >= 4 && strings.ToUpper(input[:4]) == "MAIL"
			},
			onMail,
		},
		&handler{
			func(input string) bool {
				return len(input) >= 4 && strings.ToUpper(input[:4]) == "RCPT"
			},
			onRcpt,
		},
		&handler{
			func(input string) bool {
				return len(input) >= 4 && strings.ToUpper(input[:4]) == "DATA"
			},
			onData,
		},
		&handler{
			func(input string) bool {
				return len(input) >= 4 && strings.ToUpper(input[:4]) == "RSET"
			},
			onRset,
		},
		&handler{
			func(input string) bool {
				return len(input) >= 4 && strings.ToUpper(input[:4]) == "NOOP"
			},
			onNoop,
		},
		&handler{
			func(input string) bool {
				return len(input) >= 4 && strings.ToUpper(input[:4]) == "VRFY"
			},
			onVrfy,
		},
		&handler{
			func(input string) bool {
				return len(input) >= 4 && strings.ToUpper(input[:4]) == "QUIT"
			},
			onQuit,
		},
	)
}

func onHelo(message string, conn *connection) {
	parts := strings.Fields(message)[1:]
	if len(parts) == 0 {
		conn.mustSend(501, "HELO requires domain/address - see RFC-5321 4.1.1.1")
		return
	}
	conn.session.changeState(StateHelo)
	conn.mustSend(250, fmt.Sprintf("-%s greets %s", conn.Server.HostName, parts[0]))
}

func onEhlo(message string, conn *connection) {
	parts := strings.Fields(message)[1:]
	if len(parts) == 0 {
		conn.mustSend(501, "EHLO requires domain/address - see RFC-5321 4.1.1.1")
		return
	}
	conn.session.changeState(StateEhlo)
	conn.mustSend(250, fmt.Sprintf("-%s greets %s", conn.Server.HostName, parts[0]))
}

func onMail(message string, conn *connection) {
	re := regexp.MustCompile(`FROM:\<(\S+@\S+)\>`)
	matches := re.FindStringSubmatch(strings.Trim(message, "\r"))
	if len(matches) < 1 {
		// TODO handle better.
		panic("NOT ENOUGH MATCHES")
	}

	// TODO do something with from
	//from := matches[1]

	conn.session.changeState(StateMail)
	conn.mustSend(250, "OK")
}

func onRcpt(message string, conn *connection) {
	// Check To
	conn.session.changeState(StateRcpt)
	conn.mustSend(220, message)
}

func onData(_ string, conn *connection) {
	conn.mustSend(354, "Start mail input; end with <CRLF>.<CRLF>")
	// Loop on receipt, look for <crlf>.<crlf>
	conn.session.changeState(StateData)
	conn.mustSend(250, "OK")
}

func onRset(message string, conn *connection) {
	// TODO
	conn.session.changeState(StateRset)
	conn.mustSend(220, message)
}

func onNoop(message string, conn *connection) {
	// TODO
	conn.session.changeState(StateNoop)
	conn.mustSend(220, message)
}

func onVrfy(message string, conn *connection) {
	// TODO
	conn.session.changeState(StateVrfy)
	conn.mustSend(220, message)
}

func onQuit(_ string, conn *connection) {
	conn.session.changeState(StateQuit)
	conn.mustSend(221, fmt.Sprintf("%s closing connection.", conn.Server.HostName))
}

func handleClientInput(input string, conn *connection) {
	for _, handler := range handlers {
		if handler.match(input) {
			handler.clientInputHandler(input, conn)
			return
		}
	}

	// TODO Setup a map of standard errors codes.
	conn.mustSend(500, "Unrecognized command")
}

const (
	StateInit State = "init"
	StateEhlo State = "ehlo"
	StateHelo State = "helo"
	StateMail State = "mail"
	StateRcpt State = "rcpt"
	StateData State = "data"
	StateRset State = "rset"
	StateNoop State = "noop"
	StateVrfy State = "vrfy"
	StateQuit State = "quit"
)
