package smtp

import (
	"fmt"
	"log"
)

type State string
type event string

// An FSM representing the states of an SMTP interaction.
type Smtp struct {
	State State
}

// Constructor for the SMTP FSM.
func NewSmtp() Smtp {
	return Smtp{
		StateInit,
	}
}

type transitionError struct {
	s State
	e event
}

func (e *transitionError) Error() string {
	return fmt.Sprintf("Current state %s does not understand event %s", e.s, e.e)
}

// Transition Smtp using an event.
// If the passed event isn't valid for the current state,
// an error is returned. This is a logical error
// due to an invalid event being passed.
func (smtp *Smtp) Transition(e event) error {
	if endStatesList.contains(smtp.State) {
		return &transitionError{smtp.State, e}
	}
	stateTransitions, ok := transitions[smtp.State]
	if !ok {
		// Only here if not all states were registered with the transitions map.
		log.Fatalf("State %s not registered", smtp.State)
	}
	newState, ok := stateTransitions[e]
	if !ok {
		// Cannot apply this event to the current state.
		return &transitionError{smtp.State, e}
	}
	smtp.State = newState

	return nil
}

func (smtp *Smtp) Is(s State) bool {
	return smtp.State == s
}

//func (fsm *StateMachine) Cannot(s State) bool {
//	return fsm.S.Cannot(string(s))
//}
//
//func (fsm *StateMachine) IsFinished() bool {
//	return fsm.S.Is(string(StateComp))
//}
//
//func (fsm *StateMachine) stateResponse() response {
//	return stateResponses[State(fsm.S.Current())]
//}
//
//func (fsm *StateMachine) setState(s State) {
//	fsm.S.SetState(string(s))
//}
//
//
//func NewSmtpStateMachine() *StateMachine {
//	smtpFsm := &StateMachine{}
//
//	events := fsm.Events{
//		{Name: string(eventHelo), Src: []string{string(StateInit)}, Dst: string(StateHelo)},
//		{Name: string(eventEhlo), Src: []string{string(StateInit)}, Dst: string(StateHelo)},
//		{Name: string(eventMail), Src: []string{string(StateHelo)}, Dst: string(StateMail)},
//		{Name: string(eventRcpt), Src: []string{string(StateMail)}, Dst: string(StateRcpt)},
//		{Name: string(eventData), Src: []string{string(StateRcpt)}, Dst: string(StateData)},
//
//		{Name: string(eventRset), Src: []string{string(StateMail), string(StateRcpt)}, Dst: string(StateHelo)},
//		{Name: string(eventPnnd), Src: []string{string(StateData)}, Dst: string(StateHelo)},
//		{Name: string(eventQuit), Src: []string{string(StateInit), string(StateHelo), string(StateMail), string(StateRcpt)}, Dst: string(StateComp)},
//	}
//
//	smtpFsm.S = fsm.NewFSM(string(StateInit), events, fsm.Callbacks{})
//
//	return smtpFsm
//}
//
//type response struct {
//	code    int
//	message string
//}
//
//func (res *response) string() string {
//	return fmt.Sprintf("%d %s", res.code, res.message)
//}
//
//func (res *response) bytes() []byte {
//	return []byte(res.string())
//}
//
//var stateResponses = map[State]response{
//	StateInit: {220, "Mailmole smtp"},
//	StateHelo: {250, "Mailmole at your service"},
//	StateMail: {250, "Mailmole at your service"},
//	State: {250, "Mailmole at your service"},
//}

const (
	StateInit State = "init"
	StateHelo State = "helo"
	StateEhlo State = "ehlo"
	StateMail State = "mail"
	StateRcpt State = "rcpt"
	StateData State = "data"
	StateQuit State = "quit"
	StateComp State = "complete"

	eventHelo event = "helo"
	eventEhlo event = "ehlo"
	eventMail event = "mail"
	eventRcpt event = "rcpt"
	eventData event = "data"
	eventRset event = "reset"
	eventPnnd event = "pound"
	eventSend event = "send"
	eventQuit event = "quit"
)

var transitions map[State]map[event]State
var endStatesList endStates
type endStates []State

func (endStates *endStates) contains (s State) bool {
	for _, state := range *endStates {
		if state == s {
			return true
		}
	}

	return false
}

func init() {
	endStatesList = []State{
		StateComp,
		StateQuit,
	}

	transitions = make(map[State]map[event]State)

	transitions[StateInit] = make(map[event]State)
	transitions[StateHelo] = make(map[event]State)
	transitions[StateEhlo] = make(map[event]State)
	transitions[StateMail] = make(map[event]State)
	transitions[StateRcpt] = make(map[event]State)
	transitions[StateData] = make(map[event]State)
	transitions[StateQuit] = make(map[event]State)
	transitions[StateComp] = make(map[event]State)

	transitions[StateInit][eventHelo] = StateHelo
	transitions[StateInit][eventEhlo] = StateEhlo

	transitions[StateHelo][eventMail] = StateMail
	transitions[StateHelo][eventQuit] = StateQuit
	transitions[StateHelo][eventRset] = StateHelo

	transitions[StateEhlo][eventMail] = StateMail
	transitions[StateEhlo][eventQuit] = StateQuit
	transitions[StateEhlo][eventRset] = StateEhlo

	transitions[StateMail][eventRcpt] = StateRcpt
	transitions[StateMail][eventQuit] = StateQuit
	transitions[StateMail][eventRset] = StateHelo // Not sure if this (and others) are suitable? Do we need to reset to Ehlo if that was used?

	transitions[StateRcpt][eventRcpt] = StateRcpt
	transitions[StateRcpt][eventQuit] = StateQuit
	transitions[StateRcpt][eventRset] = StateHelo // And again.

	transitions[StateData][eventSend] = StateComp
}
