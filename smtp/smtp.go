package smtp

import (
	"fmt"
	"github.com/looplab/fsm"
)

type State string
type event string

const (
	StateInit State = "init"
	StateHelo State = "helo"
	StateMail State = "mail"
	StateRcpt State = "rcpt"
	StateData State = "data"
	StateFini State = "fini"

	eventHelo event = "helo"
	eventEhlo       = eventHelo
	eventMail event = "mail"
	eventRcpt event = "rcpt"
	eventData event = "data"
	eventRset event = "reset"
	eventPnnd event = "pound"
	eventQuit event = "pound"
)

// A finite state machine representing the states of an SMTP interaction.
type StateMachine struct {
	S   *fsm.FSM
}

func (fsm *StateMachine) State() State {
	return State(fsm.S.Current())
}

func (fsm *StateMachine) Cannot(s State) bool {
	return fsm.S.Cannot(string(s))
}

func (fsm *StateMachine) IsFinished() bool {
	return fsm.S.Is(string(StateFini))
}

func (fsm *StateMachine) stateResponse() response {
	return stateResponses[State(fsm.S.Current())]
}

func (fsm *StateMachine) setState(s State) {
	fsm.S.SetState(string(s))
}


func NewSmtpStateMachine() *StateMachine {
	smtpFsm := &StateMachine{}

	events := fsm.Events{
		{Name: string(eventHelo), Src: []string{string(StateInit)}, Dst: string(StateHelo)},
		{Name: string(eventEhlo), Src: []string{string(StateInit)}, Dst: string(StateHelo)},
		{Name: string(eventMail), Src: []string{string(StateHelo)}, Dst: string(StateMail)},
		{Name: string(eventRcpt), Src: []string{string(StateMail)}, Dst: string(StateRcpt)},
		{Name: string(eventData), Src: []string{string(StateRcpt)}, Dst: string(StateData)},

		{Name: string(eventRset), Src: []string{string(StateMail), string(StateRcpt)}, Dst: string(StateHelo)},
		{Name: string(eventPnnd), Src: []string{string(StateData)}, Dst: string(StateHelo)},
		{Name: string(eventQuit), Src: []string{string(StateInit), string(StateHelo), string(StateMail), string(StateRcpt)}, Dst: string(StateFini)},
	}

	smtpFsm.S = fsm.NewFSM(string(StateInit), events, fsm.Callbacks{})

	return smtpFsm
}

type response struct {
	code    int
	message string
}

func (res *response) string() string {
	return fmt.Sprintf("%d %s", res.code, res.message)
}

func (res *response) bytes() []byte {
	return []byte(res.string())
}

var stateResponses = map[State]response{
	StateInit: {220, "Mailmole smtp"},
	StateHelo: {250, "Mailmole at your service"},
	StateMail: {250, "Mailmole at your service"},
	State: {250, "Mailmole at your service"},
}
